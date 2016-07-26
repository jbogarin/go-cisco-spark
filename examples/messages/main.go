package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	sparkClient := ciscospark.NewClient(client)
	token := "" // Change to your token
	sparkClient.Authorization = "Bearer " + token

	/*

		MESSAGES

	*/

	myRoomID := "" // Change to your testing room

	// POST messages - Text Message

	message := &ciscospark.MessageRequest{
		Text:   "This is a text message",
		RoomID: myRoomID,
	}
	newTextMessage, _, err := sparkClient.Messages.Post(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST:", newTextMessage.ID, newTextMessage.Text, newTextMessage.Created)

	// POST messages - Markdown Message

	markDownMessage := &ciscospark.MessageRequest{
		MarkDown: "This is a markdown message. *Italic*, **bold** and ***italic/bold***.",
		RoomID:   myRoomID,
	}
	newMarkDownMessage, _, err := sparkClient.Messages.Post(markDownMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST:", newMarkDownMessage.ID, newMarkDownMessage.MarkDown, newMarkDownMessage.Created)

	// POST messages - Markdown Message

	htmlMessage := &ciscospark.MessageRequest{
		MarkDown: "This is a html message with <strong>strong</strong>",
		RoomID:   myRoomID,
	}
	newHTMLMessage, _, err := sparkClient.Messages.Post(htmlMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST:", newHTMLMessage.ID, newHTMLMessage.MarkDown, newHTMLMessage.Created)

	// GET messages
	messageQueryParams := &ciscospark.MessageQueryParams{
		Max:    5,
		RoomID: myRoomID,
	}

	messages, _, err := sparkClient.Messages.Get(messageQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, message := range messages {
		fmt.Println("GET:", id, message.ID, message.Text, message.Created)
	}

	// GET messages/<ID>

	htmlMessageGet, _, err := sparkClient.Messages.GetMessage(newHTMLMessage.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", htmlMessageGet.ID, htmlMessageGet.Text, htmlMessageGet.Created)

	resp, err := sparkClient.Messages.DeleteMessage(newTextMessage.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DELETE:", resp.StatusCode)

}
