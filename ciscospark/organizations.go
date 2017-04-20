package ciscospark

const organizationsBasePath = "v1/organizations"

// organizationsService is an interface for interfacing with the organizations
// endpoints of the Cisco Spark API
type OrganizationsService service

// GetOrganizationsQueryParams ...
type GetOrganizationsQueryParams struct {
	Max int `url:"max,omitempty"`
}

// Organization represents the Spark organizations
type Organization struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Created     string `json:"created,omitempty"`
}

type organizationsRoot struct {
	Organizations []*Organization `json:"items"`
}

func (r Organization) String() string {
	return Stringify(r)
}

// Get ....
func (s *OrganizationsService) Get(queryParams *GetOrganizationsQueryParams) ([]*Organization, *Response, error) {
	path := organizationsBasePath
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(organizationsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Organizations, resp, err
}

// GetOrganization ....
func (s *OrganizationsService) GetOrganization(OrganizationID string) (*Organization, *Response, error) {
	path := organizationsBasePath + "/" + OrganizationID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	Organization := new(Organization)
	resp, err := s.client.Do(req, Organization)
	if err != nil {
		return nil, resp, err
	}

	return Organization, resp, err
}
