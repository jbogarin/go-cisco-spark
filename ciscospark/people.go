package ciscospark

const peopleBasePath = "v1/people"

// PeopleService is an interface for interfacing with the People
// endpoints of the Cisco Spark API
type PeopleService interface {
	Post(*PersonRequest) (*Person, *Response, error)
	Get(*PersonQueryParams) ([]*Person, *Response, error)
	GetPerson(string) (*Person, *Response, error)
	DeletePerson(string) (*Response, error)
	UpdatePerson(string, *UpdatePersonRequest) (*Person, *Response, error)
}

// PeopleServiceOp handles communication with the People related methods of
// the Cisco Spark API.
type PeopleServiceOp struct {
	client *Client
}

var _ PeopleService = &PeopleServiceOp{}

// PersonQueryParams ...
type PersonQueryParams struct {
	Max int `url:"max,omitempty"`
}

// PersonRequest represents the Spark people
type PersonRequest struct {
	Name      string `json:"name,omitempty"`
	TargetURL string `json:"targetUrl,omitempty"`
	Resource  string `json:"resource,omitempty"`
	Event     string `json:"event,omitempty"`
	Filter    string `json:"filter,omitempty"`
}

// UpdatePersonRequest represents the Spark people
type UpdatePersonRequest struct {
	Name      string `json:"name,omitempty"`
	TargetURL string `json:"targetUrl,omitempty"`
}

// Person ...
type Person struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	TargetURL string `json:"targetUrl,omitempty"`
	Resource  string `json:"resource,omitempty"`
	Event     string `json:"event,omitempty"`
	Filter    string `json:"filter,omitempty"`
	Created   string `json:"created,omitempty"`
}

type peopleRoot struct {
	People []*Person `json:"items"`
}

func (r Person) String() string {
	return Stringify(r)
}

func (r PersonRequest) String() string {
	return Stringify(r)
}

// Get ....
func (s *PeopleServiceOp) Get(queryParams *PersonQueryParams) ([]*Person, *Response, error) {
	path := peopleBasePath
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(peopleRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.People, resp, err

}

// Post ....
func (s *PeopleServiceOp) Post(webhookRequest *PersonRequest) (*Person, *Response, error) {
	path := peopleBasePath

	req, err := s.client.NewRequest("POST", path, webhookRequest)
	if err != nil {
		return nil, nil, err
	}

	response := new(Person)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, err
}

// GetPerson ....
func (s *PeopleServiceOp) GetPerson(webhookID string) (*Person, *Response, error) {
	path := peopleBasePath + "/" + webhookID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	webhook := new(Person)
	resp, err := s.client.Do(req, webhook)
	if err != nil {
		return nil, resp, err
	}

	return webhook, resp, err
}

// UpdatePerson ....
func (s *PeopleServiceOp) UpdatePerson(webhookID string, updatePersonRequest *UpdatePersonRequest) (*Person, *Response, error) {
	path := peopleBasePath + "/" + webhookID

	req, err := s.client.NewRequest("PUT", path, updatePersonRequest)
	if err != nil {
		return nil, nil, err
	}

	webhook := new(Person)
	resp, err := s.client.Do(req, webhook)
	if err != nil {
		return nil, resp, err
	}

	return webhook, resp, err
}

// DeletePerson ....
func (s *PeopleServiceOp) DeletePerson(webhookID string) (*Response, error) {
	path := peopleBasePath + "/" + webhookID

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
