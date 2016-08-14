package ciscospark

const teamMembershipsBasePath = "v1/team/memberships"

// TeamMembershipsService is an interface for interfacing with the TeamMemberships
// endpoints of the Cisco Spark API
type TeamMembershipsService service

// TeamMembershipQueryParams ...
type TeamMembershipQueryParams struct {
	Max    int    `url:"max,omitempty"`
	TeamID string `url:"teamId,omitempty"`
}

// TeamMembershipRequest represents the Spark teamMemberships
type TeamMembershipRequest struct {
	TeamID      string `json:"teamId,omitempty"`
	PersonID    string `json:"personId,omitempty"`
	PersonEmail string `json:"personEmail,omitempty"`
	IsModerator bool   `json:"isModerator,omitempty"`
}

// UpdateTeamMembershipRequest represents the Spark teamMemberships
type UpdateTeamMembershipRequest struct {
	IsModerator bool `json:"isModerator,omitempty"`
}

// TeamMembership ...
type TeamMembership struct {
	ID          string `json:"id,omitempty"`
	TeamID      string `json:"teamId,omitempty"`
	PersonID    string `json:"personId,omitempty"`
	PersonEmail string `json:"personEmail,omitempty"`
	IsModerator bool   `json:"isModerator,omitempty"`
	Created     string `json:"created,omitempty"`
}

type teamMembershipsRoot struct {
	TeamMemberships []*TeamMembership `json:"items"`
}

func (r TeamMembership) String() string {
	return Stringify(r)
}

func (r TeamMembershipRequest) String() string {
	return Stringify(r)
}

// Get ....
func (s *TeamMembershipsService) Get(queryParams *TeamMembershipQueryParams) ([]*TeamMembership, *Response, error) {
	path := teamMembershipsBasePath
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(teamMembershipsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.TeamMemberships, resp, err

}

// Post ....
func (s *TeamMembershipsService) Post(teamRequest *TeamMembershipRequest) (*TeamMembership, *Response, error) {
	path := teamMembershipsBasePath

	req, err := s.client.NewRequest("POST", path, teamRequest)
	if err != nil {
		return nil, nil, err
	}

	response := new(TeamMembership)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, err
}

// GetTeamMembership ....
func (s *TeamMembershipsService) GetTeamMembership(teamID string) (*TeamMembership, *Response, error) {
	path := teamMembershipsBasePath + "/" + teamID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	team := new(TeamMembership)
	resp, err := s.client.Do(req, team)
	if err != nil {
		return nil, resp, err
	}

	return team, resp, err
}

// UpdateTeamMembership ....
func (s *TeamMembershipsService) UpdateTeamMembership(teamID string, updateTeamMembershipRequest *UpdateTeamMembershipRequest) (*TeamMembership, *Response, error) {
	path := teamMembershipsBasePath + "/" + teamID

	req, err := s.client.NewRequest("PUT", path, updateTeamMembershipRequest)
	if err != nil {
		return nil, nil, err
	}

	team := new(TeamMembership)
	resp, err := s.client.Do(req, team)
	if err != nil {
		return nil, resp, err
	}

	return team, resp, err
}

// DeleteTeamMembership ....
func (s *TeamMembershipsService) DeleteTeamMembership(teamID string) (*Response, error) {
	path := teamMembershipsBasePath + "/" + teamID

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
