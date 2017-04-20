package ciscospark

const peopleBasePath = "v1/people"

// PeopleService is an interface for interfacing with the People
// endpoints of the Cisco Spark API
type PeopleService service

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
	FirstName   string   `json:"firstName,omitempty"`
	LastName    string   `json:"lastName,omitempty"`
	Avatar      string   `json:"avatar,omitempty"`
	Created     string   `json:"created,omitempty"`
	OrgID       string   `json:"orgId,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Licenses    []string `json:"licenses,omitempty"`
	TimeZone    string   `json:"timezone,omitempty"`
	Status      string   `json:"status,omitempty"`
	NickName    string   `json:"nickName,omitempty"`
	Type        string   `json:"type,omitempty"`
}

type peopleRoot struct {
	People []*Person `json:"items"`
}

func (r Person) String() string {
	return Stringify(r)
}

// Get ....
func (s *PeopleService) Get(queryParams *GetPeopleQueryParams) ([]*Person, *Response, error) {
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
func (s *PeopleService) GetPerson(personID string) (*Person, *Response, error) {
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
func (s *PeopleService) GetMe() (*Person, *Response, error) {
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
