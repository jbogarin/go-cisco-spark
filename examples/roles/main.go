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

		Roles

	*/

	// GET Roles
	queryParams := &ciscospark.GetRolesQueryParams{
		Max: 2,
	}

	Roles, _, err := sparkClient.Roles.Get(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	RoleID := ""
	for id, Role := range Roles {
		fmt.Println("GET:", id, Role.ID, Role.Name)
		RoleID = Role.ID
	}

	// GET Roles/<id>
	Role, _, err := sparkClient.Roles.GetRole(RoleID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", Role.ID, Role.Name)

}
