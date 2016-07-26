package ciscospark

const messagesBasePath = "v1/messages"

// MessagesService is an interface for interfacing with the Messages
// endpoints of the Cisco Spark API
type MessagesService interface {
	Post(*MessageRequest) (*Message, *Response, error)
	Get(*MessageQueryParams) ([]*Message, *Response, error)
	GetMessage(string) (*Message, *Response, error)
	DeleteMessage(string) (*Response, error)
}

// MessagesServiceOp handles communication with the Messages related methods of
// the Cisco Spark API.
type MessagesServiceOp struct {
	client *Client
}

var _ MessagesService = &MessagesServiceOp{}

// MessageQueryParams ...
type MessageQueryParams struct {
	RoomID        string `url:"roomId,omitempty"`
	Before        string `url:"before,omitempty"`
	BeforeMessage string `url:"beforeMessage,omitempty"`
	Max           int    `url:"max,omitempty"`
}

// MessageRequest represents the Spark messages
type MessageRequest struct {
	RoomID        string   `json:"roomId,omitempty"`
	Text          string   `json:"text,omitempty"`
	Files         []string `json:"files,omitempty"`
	ToPersonID    string   `json:"toPersonId,omitempty"`
	ToPersonEmail string   `json:"toPersonEmail,omitempty"`
	MarkDown      string   `json:"markdown,omitempty"`
	HTML          string   `json:"html,omitempty"`
}

// Message ...
type Message struct {
	ID            string   `json:"id,omitempty"`
	RoomID        string   `json:"roomId,omitempty"`
	ToPersonEmail string   `json:"toPersonEmail,omitempty"`
	ToPersonID    string   `json:"toPersonId,omitempty"`
	Text          string   `json:"text,omitempty"`
	PersonID      string   `json:"personId,omitempty"`
	PersonEmail   string   `json:"personEmail,omitempty"`
	Created       string   `json:"created,omitempty"`
	MarkDown      string   `json:"markdown,omitempty"`
	HTML          string   `json:"html,omitempty"`
	Files         []string `json:"files,omitempty"`
	RoomType      string   `json:"roomType,omitempty"`
}

type messagesRoot struct {
	Messages []*Message `json:"items"`
}

func (r Message) String() string {
	return Stringify(r)
}

func (r MessageRequest) String() string {
	return Stringify(r)
}

// Get ....
func (s *MessagesServiceOp) Get(queryParams *MessageQueryParams) ([]*Message, *Response, error) {
	path := messagesBasePath
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(messagesRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Messages, resp, err

}

// Post ....
func (s *MessagesServiceOp) Post(messageRequest *MessageRequest) (*Message, *Response, error) {
	path := messagesBasePath

	req, err := s.client.NewRequest("POST", path, messageRequest)
	if err != nil {
		return nil, nil, err
	}

	response := new(Message)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, err
}

// GetMessage ....
func (s *MessagesServiceOp) GetMessage(messageID string) (*Message, *Response, error) {
	path := messagesBasePath + "/" + messageID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	message := new(Message)
	resp, err := s.client.Do(req, message)
	if err != nil {
		return nil, resp, err
	}

	return message, resp, err
}

// DeleteMessage ....
func (s *MessagesServiceOp) DeleteMessage(messageID string) (*Response, error) {
	path := messagesBasePath + "/" + messageID

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
