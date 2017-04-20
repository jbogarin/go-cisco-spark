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

		TEAM MEMBERSHIPS

	*/

	myTeamID := ""      // Change to your test team
	newTeamMember := "" // Change to the person email you want to add to the team

	// POST team-memberships
	teamMembershipRequest := &ciscospark.TeamMembershipRequest{
		TeamID:      myTeamID,
		PersonEmail: newTeamMember,
		IsModerator: true,
	}

	newTeamMembership, _, err := sparkClient.TeamMemberships.Post(teamMembershipRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", newTeamMembership.ID, newTeamMembership.PersonEmail, newTeamMembership.IsModerator, newTeamMembership.Created)

	// GET team-memberships
	teamMembershipsQueryParams := &ciscospark.TeamMembershipQueryParams{
		Max:    2,
		TeamID: myTeamID,
	}

	teamMemberships, _, err := sparkClient.TeamMemberships.Get(teamMembershipsQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, teamMembership := range teamMemberships {
		fmt.Println("GET:", id, teamMembership.ID, teamMembership.PersonEmail, teamMembership.IsModerator, teamMembership.Created)
	}

	// GET team-memberships/<id>
	teamMembership, _, err := sparkClient.TeamMemberships.GetTeamMembership(newTeamMembership.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET <ID>:", teamMembership.ID, teamMembership.PersonEmail, teamMembership.IsModerator, teamMembership.Created)

	// PUT team-memberships/<id>
	updateTeamMembershipRequest := &ciscospark.UpdateTeamMembershipRequest{
		IsModerator: false,
	}

	updatedTeamMembership, _, err := sparkClient.TeamMemberships.UpdateTeamMembership(newTeamMembership.ID, updateTeamMembershipRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedTeamMembership.ID, updatedTeamMembership.PersonEmail, updatedTeamMembership.IsModerator, updatedTeamMembership.Created)

	// DELETE team-memberships/<id>
	resp, err := sparkClient.TeamMemberships.DeleteTeamMembership(newTeamMembership.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", resp.StatusCode)

}
