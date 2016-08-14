package ciscospark

const membershipsBasePath = "v1/memberships"

// MembershipsService handles communication with the Memberships related methods of
// the Cisco Spark API.
type MembershipsService service

// MembershipQueryParams ...
type MembershipQueryParams struct {
	RoomID      string `url:"roomId,omitempty"`
	PersonID    string `url:"personId,omitempty"`
	PersonEmail string `url:"personEmail,omitempty"`
	Max         int    `url:"max,omitempty"`
}

// MembershipRequest represents the Spark memberships
type MembershipRequest struct {
	RoomID      string `json:"roomId,omitempty"`
	PersonID    string `json:"personId,omitempty"`
	PersonEmail string `json:"personEmail,omitempty"`
	IsModerator bool   `json:"isModerator,omitempty"`
}

// UpdateMembershipRequest represents the Spark memberships
type UpdateMembershipRequest struct {
	IsModerator bool `json:"isModerator,omitempty"`
}

// Membership ...
type Membership struct {
	ID          string `json:"id,omitempty"`
	RoomID      string `json:"roomId,omitempty"`
	PersonID    string `json:"personId,omitempty"`
	PersonEmail string `json:"personEmail,omitempty"`
	Created     string `json:"created,omitempty"`
	IsModerator bool   `json:"isModerator,omitempty"`
	IsMonitor   bool   `json:"isMonitor,omitempty"`
}

type membershipsRoot struct {
	Memberships []*Membership `json:"items"`
}

func (r Membership) String() string {
	return Stringify(r)
}

func (r MembershipRequest) String() string {
	return Stringify(r)
}

// Get ....
func (s *MembershipsService) Get(queryParams *MembershipQueryParams) ([]*Membership, *Response, error) {
	path := membershipsBasePath
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(membershipsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Memberships, resp, err

}

// Post ....
func (s *MembershipsService) Post(membershipRequest *MembershipRequest) (*Membership, *Response, error) {
	path := membershipsBasePath

	req, err := s.client.NewRequest("POST", path, membershipRequest)
	if err != nil {
		return nil, nil, err
	}

	response := new(Membership)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, err
}

// GetMembership ....
func (s *MembershipsService) GetMembership(membershipID string) (*Membership, *Response, error) {
	path := membershipsBasePath + "/" + membershipID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	membership := new(Membership)
	resp, err := s.client.Do(req, membership)
	if err != nil {
		return nil, resp, err
	}

	return membership, resp, err
}

// UpdateMembership ....
func (s *MembershipsService) UpdateMembership(membershipID string, updateMembershipRequest *UpdateMembershipRequest) (*Membership, *Response, error) {
	path := membershipsBasePath + "/" + membershipID

	req, err := s.client.NewRequest("PUT", path, updateMembershipRequest)
	if err != nil {
		return nil, nil, err
	}

	membership := new(Membership)
	resp, err := s.client.Do(req, membership)
	if err != nil {
		return nil, resp, err
	}

	return membership, resp, err
}

// DeleteMembership ....
func (s *MembershipsService) DeleteMembership(membershipID string) (*Response, error) {
	path := membershipsBasePath + "/" + membershipID

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
