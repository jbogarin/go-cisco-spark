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

		Licenses

	*/

	// GET Licenses
	queryParams := &ciscospark.GetLicensesQueryParams{
		Max: 2,
	}

	Licenses, _, err := sparkClient.Licenses.Get(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	LicenseID := ""
	for id, License := range Licenses {
		fmt.Println("GET:", id, License.ID, License.Name, License.TotalUnits, License.ConsumedUnits)
		LicenseID = License.ID
	}

	// GET Licenses/<id>
	License, _, err := sparkClient.Licenses.GetLicense(LicenseID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", License.ID, License.Name, License.TotalUnits, License.ConsumedUnits)

}
