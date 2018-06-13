package main

import (
	"log"
	"flag"
	"os"
	"fmt"
)

func showUsage() {
	fmt.Fprintf(os.Stderr, "./%s\n", flag.CommandLine.Name())
	flag.CommandLine.PrintDefaults()
}

func main() {
	topicId := flag.Int("topicId", -1, "a topic id")
	token := flag.String("token", "", "Typetalk bot token")
	url := flag.String("url", "https://typetalk.com", "Typetalk url")
	flag.Parse()

	rest := flag.Args()
	if (*topicId == -1 || *token == "" || len(rest) == 0) {
		showUsage()
	} else {
		log.Println("Starting typetalk cli...")
		cli := NewCli(*url, *token)
		err, _ := cli.PostMessage(*topicId, rest[0], rest[1:])
		if (err != nil) {
			log.Fatal(err)
		} else {
			log.Println("Completed")
		}
	}

}
