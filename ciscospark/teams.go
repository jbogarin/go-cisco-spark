package ciscospark

const teamsBasePath = "v1/teams"

// TeamsService is an interface for interfacing with the Teams
// endpoints of the Cisco Spark API
type TeamsService interface {
	Post(*TeamRequest) (*Team, *Response, error)
	Get(*TeamQueryParams) ([]*Team, *Response, error)
	GetTeam(string) (*Team, *Response, error)
	DeleteTeam(string) (*Response, error)
	UpdateTeam(string, *UpdateTeamRequest) (*Team, *Response, error)
}

// TeamsServiceOp handles communication with the Teams related methods of
// the Cisco Spark API.
type TeamsServiceOp struct {
	client *Client
}

var _ TeamsService = &TeamsServiceOp{}

// TeamQueryParams ...
type TeamQueryParams struct {
	Max int `url:"max,omitempty"`
}

// TeamRequest represents the Spark teams
type TeamRequest struct {
	Name string `json:"name,omitempty"`
}

// UpdateTeamRequest represents the Spark teams
type UpdateTeamRequest struct {
	Name string `json:"name,omitempty"`
}

// Team ...
type Team struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
}

type teamsRoot struct {
	Teams []*Team `json:"items"`
}

func (r Team) String() string {
	return Stringify(r)
}

func (r TeamRequest) String() string {
	return Stringify(r)
}

// Get ....
func (s *TeamsServiceOp) Get(queryParams *TeamQueryParams) ([]*Team, *Response, error) {
	path := teamsBasePath
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(teamsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Teams, resp, err

}

// Post ....
func (s *TeamsServiceOp) Post(teamRequest *TeamRequest) (*Team, *Response, error) {
	path := teamsBasePath

	req, err := s.client.NewRequest("POST", path, teamRequest)
	if err != nil {
		return nil, nil, err
	}

	response := new(Team)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, err
}

// GetTeam ....
func (s *TeamsServiceOp) GetTeam(teamID string) (*Team, *Response, error) {
	path := teamsBasePath + "/" + teamID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	team := new(Team)
	resp, err := s.client.Do(req, team)
	if err != nil {
		return nil, resp, err
	}

	return team, resp, err
}

// UpdateTeam ....
func (s *TeamsServiceOp) UpdateTeam(teamID string, updateTeamRequest *UpdateTeamRequest) (*Team, *Response, error) {
	path := teamsBasePath + "/" + teamID

	req, err := s.client.NewRequest("PUT", path, updateTeamRequest)
	if err != nil {
		return nil, nil, err
	}

	team := new(Team)
	resp, err := s.client.Do(req, team)
	if err != nil {
		return nil, resp, err
	}

	return team, resp, err
}

// DeleteTeam ....
func (s *TeamsServiceOp) DeleteTeam(teamID string) (*Response, error) {
	path := teamsBasePath + "/" + teamID

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
