package ciscospark

const webhooksBasePath = "v1/webhooks"

// WebhooksService is an interface for interfacing with the Webhooks
// endpoints of the Cisco Spark API
type WebhooksService service

// WebhookQueryParams ...
type WebhookQueryParams struct {
	Max int `url:"max,omitempty"`
}

// WebhookRequest represents the Spark webhooks
type WebhookRequest struct {
	Name      string `json:"name,omitempty"`
	TargetURL string `json:"targetUrl,omitempty"`
	Resource  string `json:"resource,omitempty"`
	Event     string `json:"event,omitempty"`
	Filter    string `json:"filter,omitempty"`
}

// UpdateWebhookRequest represents the Spark webhooks
type UpdateWebhookRequest struct {
	Name      string `json:"name,omitempty"`
	TargetURL string `json:"targetUrl,omitempty"`
}

// Webhook ...
type Webhook struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	TargetURL string `json:"targetUrl,omitempty"`
	Resource  string `json:"resource,omitempty"`
	Event     string `json:"event,omitempty"`
	Filter    string `json:"filter,omitempty"`
	Created   string `json:"created,omitempty"`
}

type webhooksRoot struct {
	Webhooks []*Webhook `json:"items"`
}

func (r Webhook) String() string {
	return Stringify(r)
}

func (r WebhookRequest) String() string {
	return Stringify(r)
}

// Get ....
func (s *WebhooksService) Get(queryParams *WebhookQueryParams) ([]*Webhook, *Response, error) {
	path := webhooksBasePath
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(webhooksRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Webhooks, resp, err

}

// Post ....
func (s *WebhooksService) Post(webhookRequest *WebhookRequest) (*Webhook, *Response, error) {
	path := webhooksBasePath

	req, err := s.client.NewRequest("POST", path, webhookRequest)
	if err != nil {
		return nil, nil, err
	}

	response := new(Webhook)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, err
}

// GetWebhook ....
func (s *WebhooksService) GetWebhook(webhookID string) (*Webhook, *Response, error) {
	path := webhooksBasePath + "/" + webhookID

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	webhook := new(Webhook)
	resp, err := s.client.Do(req, webhook)
	if err != nil {
		return nil, resp, err
	}

	return webhook, resp, err
}

// UpdateWebhook ....
func (s *WebhooksService) UpdateWebhook(webhookID string, updateWebhookRequest *UpdateWebhookRequest) (*Webhook, *Response, error) {
	path := webhooksBasePath + "/" + webhookID

	req, err := s.client.NewRequest("PUT", path, updateWebhookRequest)
	if err != nil {
		return nil, nil, err
	}

	webhook := new(Webhook)
	resp, err := s.client.Do(req, webhook)
	if err != nil {
		return nil, resp, err
	}

	return webhook, resp, err
}

// DeleteWebhook ....
func (s *WebhooksService) DeleteWebhook(webhookID string) (*Response, error) {
	path := webhooksBasePath + "/" + webhookID

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
