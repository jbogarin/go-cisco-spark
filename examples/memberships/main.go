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

		MEMBERSHIPS

	*/

	// This works if you create a room where you are a moderator (paid feature). I tested with a room that it is part of a team.

	myRoomID := ""       // Change to your testing room
	myTestingEmail := "" // Change to your email

	// GET memberships

	membershipQueryParams := &ciscospark.MembershipQueryParams{
		Max:         2,
		PersonEmail: myTestingEmail,
	}

	memberships, _, err := sparkClient.Memberships.Get(membershipQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for _, membership := range memberships {
		fmt.Println("GET:", membership.ID, membership.PersonEmail, membership.IsModerator, membership.Created)
	}

	// POST memberships

	membershipRequest := &ciscospark.MembershipRequest{
		RoomID:      myRoomID,
		PersonEmail: myTestingEmail,
		IsModerator: true,
	}

	testMembership, _, err := sparkClient.Memberships.Post(membershipRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", testMembership.ID, testMembership.PersonEmail, testMembership.IsModerator, testMembership.Created)

	// GET memberships/<ID>

	membership, _, err := sparkClient.Memberships.GetMembership(testMembership.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET <ID>:", membership.ID, membership.PersonEmail, membership.IsModerator, membership.Created)

	updateMembershipRequest := &ciscospark.UpdateMembershipRequest{
		IsModerator: false,
	}

	// PUT memberships/<ID>

	updatedMembership, _, err := sparkClient.Memberships.UpdateMembership(testMembership.ID, updateMembershipRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PUT:", updatedMembership.ID, updatedMembership.PersonEmail, updatedMembership.IsModerator, updatedMembership.Created)

	// DELETE memberships<ID>

	resp, err := sparkClient.Memberships.DeleteMembership(testMembership.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", resp.StatusCode)

}
