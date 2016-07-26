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

		TEAMS

	*/

	// POST teams
	teamRequest := &ciscospark.TeamRequest{
		Name: "Go Test Team",
	}

	newTeam, _, err := sparkClient.Teams.Post(teamRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", newTeam.ID, newTeam.Name, newTeam.Created)

	// GET teams

	teamQueryParams := &ciscospark.TeamQueryParams{
		Max: 2,
	}

	teams, _, err := sparkClient.Teams.Get(teamQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, team := range teams {
		fmt.Println("GET:", id, team.ID, team.Name, team.Created)
	}

	// GET teams/<id>
	team, _, err := sparkClient.Teams.GetTeam(newTeam.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET <ID>:", team.ID, team.Created, team.Name)

	// PUT teams/<id>
	updateTeamRequest := &ciscospark.UpdateTeamRequest{
		Name: "Go Test Team 2",
	}

	updatedTeam, _, err := sparkClient.Teams.UpdateTeam(newTeam.ID, updateTeamRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PUT:", updatedTeam.ID, updatedTeam.Name, updatedTeam.Created)

	// DELETE teams/<id>
	resp, err := sparkClient.Teams.DeleteTeam(newTeam.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", resp.StatusCode)

}
