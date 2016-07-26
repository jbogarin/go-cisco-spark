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

	myRoomID := ""   // Change to your testing room
	webHookURL := "" // Change this to your test URL

	// POST webhooks

	webhookRequest := &ciscospark.WebhookRequest{
		Name:      "Webhook - Test",
		TargetURL: webHookURL,
		Resource:  "messages",
		Event:     "created",
		Filter:    "roomId=" + myRoomID,
	}

	testWebhook, _, err := sparkClient.Webhooks.Post(webhookRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", testWebhook.ID, testWebhook.Name, testWebhook.TargetURL, testWebhook.Created)

	// GET webhooks

	webhooksQueryParams := &ciscospark.WebhookQueryParams{
		Max: 10,
	}

	webhooks, _, err := sparkClient.Webhooks.Get(webhooksQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, webhook := range webhooks {
		fmt.Println("GET:", id, webhook.ID, webhook.Name, webhook.TargetURL, webhook.Created)
	}

	// GET webhooks/<ID>
	webhook, _, err := sparkClient.Webhooks.GetWebhook(testWebhook.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET <ID>:", webhook.ID, webhook.Name, webhook.TargetURL, webhook.Created)

	updateWebhookRequest := &ciscospark.UpdateWebhookRequest{
		Name:      "Webhook Update - Test",
		TargetURL: webHookURL,
	}

	// PUT webhooks/<ID>
	updatedWebhook, _, err := sparkClient.Webhooks.UpdateWebhook(testWebhook.ID, updateWebhookRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PUT:", updatedWebhook.ID, updatedWebhook.Name, updatedWebhook.TargetURL, updatedWebhook.Created)

	// DELETE webhooks/<ID>
	resp, err := sparkClient.Webhooks.DeleteWebhook(testWebhook.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", resp.StatusCode)

}
