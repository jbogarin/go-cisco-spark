package ciscospark

const peopleBasePath = "v1/people"

// PeopleService is an interface for interfacing with the People
// endpoints of the Cisco Spark API
type PeopleService interface {
	Get(*GetPeopleQueryParams) ([]*Person, *Response, error)
	GetPerson(string) (*Person, *Response, error)
	GetMe() (*Person, *Response, error)
}

// PeopleServiceOp handles communication with the People related methods of
// the Cisco Spark API.
type PeopleServiceOp struct {
	client *Client
}

var _ PeopleService = &PeopleServiceOp{}

// GetPeopleQueryParams ...
type GetPeopleQueryParams struct {
	Email       string `url:"email,omitempty"`
	DisplayName string `url:"displayName,omitempty"`
	Max         int    `url:"max,omitempty"`
}

// Person represents the Spark people
type Person struct {
	ID          string   `json:"id,omitempty"`
	Emails      []string `json:"emails,omitempty"`
	DisplayName string   `json:"displayName,omitempty"`
	Avatar      string   `json:"avatar,omitempty"`
	Created     string   `json:"created,omitempty"`
}

type peopleRoot struct {
	People []*Person `json:"items"`
}

func (r Person) String() string {
	return Stringify(r)
}

// Get ....
func (s *PeopleServiceOp) Get(queryParams *GetPeopleQueryParams) ([]*Person, *Response, error) {
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

// GetPerson ....
func (s *PeopleServiceOp) GetPerson(personID string) (*Person, *Response, error) {
	path := peopleBasePath + "/" + personID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	person := new(Person)
	resp, err := s.client.Do(req, person)
	if err != nil {
		return nil, resp, err
	}

	return person, resp, err
}

// GetMe ....
func (s *PeopleServiceOp) GetMe() (*Person, *Response, error) {
	path := peopleBasePath + "/me"

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	person := new(Person)
	resp, err := s.client.Do(req, person)
	if err != nil {
		return nil, resp, err
	}

	return person, resp, err
}
