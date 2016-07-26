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

		ROOMS

	*/

	// POST rooms
	roomRequest := &ciscospark.RoomRequest{
		Title: "Go Test Room",
	}

	newRoom, _, err := sparkClient.Rooms.Post(roomRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", newRoom.ID, newRoom.Title, newRoom.IsLocked, newRoom.Created)

	// GET rooms
	roomsQueryParams := &ciscospark.RoomQueryParams{
		Max:  2,
		Type: "group",
	}

	rooms, _, err := sparkClient.Rooms.Get(roomsQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, room := range rooms {
		fmt.Println("GET:", id, room.ID, room.IsLocked, room.Title)
	}

	// GET rooms/<id>
	room, _, err := sparkClient.Rooms.GetRoom(newRoom.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET <ID>:", room.ID, room.Title, room.IsLocked, room.Created)

	updateRoomRequest := &ciscospark.UpdateRoomRequest{
		Title: "Go Test Room 2",
	}

	updatedRoom, _, err := sparkClient.Rooms.UpdateRoom(newRoom.ID, updateRoomRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PUT:", updatedRoom.ID, updatedRoom.Title, updatedRoom.IsLocked, updatedRoom.Created)

	// DELETE
	resp, err := sparkClient.Rooms.DeleteRoom(newRoom.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", resp.StatusCode)
}
