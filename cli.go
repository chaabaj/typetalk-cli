package main

import (
	"github.com/nulab/go-typetalk/typetalk/v1"
	"context"
	"os"
	"log"
)

type Cli struct {
	client *v1.Client
}

func NewCli (url string, token string) *Cli {
	client := v1.NewClient(nil)
	client.SetTypetalkToken(token)
	return &Cli{client: client}
}


func (cli * Cli) UploadFile(topicId int, path string) (error, *v1.AttachmentFile) {
	file, err := os.Open(path)
	var content []byte
	if (err != nil) {
		return err, nil
	} else {
		ctx := context.Background()
		attachment, b, err := cli.client.Files.UploadAttachmentFile(ctx, topicId, file)
		if (err != nil) {
			b.Body.Read(content)
			log.Println(content)
			return err, nil
		} else {
			return nil, attachment
		}
	}
}

func (cli * Cli) PostMessage(topicId int, msg string, paths []string) (error, *v1.PostedMessageResult) {
	ctx := context.Background()
	log.Println("Start post message")
	var attachments []string
	for i := 0; i < len(paths); i++  {
		println("Uploading file...")
		err, attachment := cli.UploadFile(topicId, paths[i])
		if (err != nil) {
			return err, nil
		} else {
			attachments = append(attachments, attachment.FileKey)
		}
	}

	msgOpts := &v1.PostMessageOptions{
		ReplyTo: 0,
		ShowLinkMeta: false,
		FileKeys: attachments,
		TalkIds: nil,
		FileUrls: nil,
		FileNames: nil,
	}
	log.Println("Posting message")
	postedMsg, _, err := cli.client.Messages.PostMessage(ctx, topicId, msg, msgOpts)
	return err, postedMsg
}