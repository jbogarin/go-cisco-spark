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

		PEOPLE

	*/

	// GET people
	queryParams := &ciscospark.GetPeopleQueryParams{
		DisplayName: "", // Change to the person you want to look for
		Max:         2,
	}

	people, _, err := sparkClient.People.Get(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	personID := ""
	for id, person := range people {
		fmt.Println("GET:", id, person.ID, person.DisplayName, person.Created)
		personID = person.ID
	}

	// GET people/<id>
	person, _, err := sparkClient.People.GetPerson(personID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", person.ID, person.DisplayName, person.Created)

	// GET people/me
	me, _, err := sparkClient.People.GetMe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET ME:", me.ID, me.DisplayName, me.Created)

}
