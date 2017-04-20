package ciscospark

const licensesBasePath = "v1/licenses"

// LicensesService is an interface for interfacing with the Licenses
// endpoints of the Cisco Spark API
type LicensesService service

// GetLicensesQueryParams ...
type GetLicensesQueryParams struct {
	Max   int    `url:"max,omitempty"`
	OrgID string `url:"orgId,omitempty"`
}

// License represents the Spark Licenses
type License struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	TotalUnits    int    `json:"totalUnits,omitempty"`
	ConsumedUnits int    `json:"consumedUnits,omitempty"`
}

type LicensesRoot struct {
	Licenses []*License `json:"items"`
}

func (r License) String() string {
	return Stringify(r)
}

// Get ....
func (s *LicensesService) Get(queryParams *GetLicensesQueryParams) ([]*License, *Response, error) {
	path := licensesBasePath
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(LicensesRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Licenses, resp, err
}

// GetLicense ....
func (s *LicensesService) GetLicense(LicenseID string) (*License, *Response, error) {
	path := licensesBasePath + "/" + LicenseID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	License := new(License)
	resp, err := s.client.Do(req, License)
	if err != nil {
		return nil, resp, err
	}

	return License, resp, err
}
