package ciscospark

const roomsBasePath = "v1/rooms"

// RoomsService is an interface for interfacing with the Rooms
// endpoints of the Cisco Spark API
type RoomsService interface {
	DeleteRoom(string) (*Response, error)
	Get(*RoomQueryParams) ([]*Room, *Response, error)
	GetRoom(string) (*Room, *Response, error)
	Post(*RoomRequest) (*Room, *Response, error)
	UpdateRoom(string, *UpdateRoomRequest) (*Room, *Response, error)
}

// RoomsServiceOp handles communication with the Rooms related methods of
// the Cisco Spark API.
type RoomsServiceOp struct {
	client *Client
}

var _ RoomsService = &RoomsServiceOp{}

// RoomQueryParams ...
type RoomQueryParams struct {
	Max    int    `url:"max,omitempty"`
	TeamID string `url:"teamId,omitempty"`
	Type   string `url:"type,omitempty"`
}

// RoomRequest represents the Spark rooms
type RoomRequest struct {
	TeamID string `json:"teamId,omitempty"`
	Title  string `json:"title,omitempty"`
}

// UpdateRoomRequest represents the Spark rooms
type UpdateRoomRequest struct {
	Title string `json:"title,omitempty"`
}

// Room ...
type Room struct {
	ID           string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Type         string `json:"roomId,omitempty"`
	IsLocked     bool   `json:"isLocked,omitempty"`
	TeamID       string `json:"teamId,omitempty"`
	LastActivity string `json:"lastActivity,omitempty"`
	Created      string `json:"created,omitempty"`
}

type roomsRoot struct {
	Rooms []*Room `json:"items"`
}

func (r Room) String() string {
	return Stringify(r)
}

func (r RoomRequest) String() string {
	return Stringify(r)
}

// Get ....
func (s *RoomsServiceOp) Get(queryParams *RoomQueryParams) ([]*Room, *Response, error) {
	path := roomsBasePath
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(roomsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Rooms, resp, err

}

// Post ....
func (s *RoomsServiceOp) Post(roomRequest *RoomRequest) (*Room, *Response, error) {
	path := roomsBasePath

	req, err := s.client.NewRequest("POST", path, roomRequest)
	if err != nil {
		return nil, nil, err
	}

	response := new(Room)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, err
}

// GetRoom ....
func (s *RoomsServiceOp) GetRoom(roomID string) (*Room, *Response, error) {
	path := roomsBasePath + "/" + roomID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	room := new(Room)
	resp, err := s.client.Do(req, room)
	if err != nil {
		return nil, resp, err
	}

	return room, resp, err
}

// UpdateRoom ....
func (s *RoomsServiceOp) UpdateRoom(roomID string, updateRoomRequest *UpdateRoomRequest) (*Room, *Response, error) {
	path := roomsBasePath + "/" + roomID

	req, err := s.client.NewRequest("PUT", path, updateRoomRequest)
	if err != nil {
		return nil, nil, err
	}

	room := new(Room)
	resp, err := s.client.Do(req, room)
	if err != nil {
		return nil, resp, err
	}

	return room, resp, err
}

// DeleteRoom ....
func (s *RoomsServiceOp) DeleteRoom(roomID string) (*Response, error) {
	path := roomsBasePath + "/" + roomID

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
