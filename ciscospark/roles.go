package ciscospark

const rolesBasePath = "v1/roles"

// rolesService is an interface for interfacing with the roles
// endpoints of the Cisco Spark API
type RolesService service

// GetRolesQueryParams ...
type GetRolesQueryParams struct {
	Max int `url:"max,omitempty"`
}

// Role represents the Spark roles
type Role struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type rolesRoot struct {
	Roles []*Role `json:"items"`
}

func (r Role) String() string {
	return Stringify(r)
}

// Get ....
func (s *RolesService) Get(queryParams *GetRolesQueryParams) ([]*Role, *Response, error) {
	path := rolesBasePath
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rolesRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Roles, resp, err
}

// GetRole ....
func (s *RolesService) GetRole(RoleID string) (*Role, *Response, error) {
	path := rolesBasePath + "/" + RoleID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	Role := new(Role)
	resp, err := s.client.Do(req, Role)
	if err != nil {
		return nil, resp, err
	}

	return Role, resp, err
}
