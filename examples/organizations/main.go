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

		Organizations

	*/

	// GET Organizations
	queryParams := &ciscospark.GetOrganizationsQueryParams{
		Max: 2,
	}

	Organizations, _, err := sparkClient.Organizations.Get(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	OrganizationID := ""
	for id, Organization := range Organizations {
		fmt.Println("GET:", id, Organization.ID, Organization.DisplayName, Organization.Created)
		OrganizationID = Organization.ID
	}

	// GET Organizations/<id>
	Organization, _, err := sparkClient.Organizations.GetOrganization(OrganizationID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", Organization.ID, Organization.DisplayName, Organization.Created)

}
