package cli

import (
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
	uri    string
}

func NewClient(uri string) Client {
	return Client{
		client: &http.Client{},
		uri:    uri,
	}
}

func (c Client) Create(title, message string, duration time.Duration) ([]byte, error) {
	res := []byte(`response for create task`)
	return res, nil
}

func (c Client) Edit(id, title, message string, duration time.Duration) ([]byte, error) {
	res := []byte(`response for edit task`)
	return res, nil
}

func (c Client) Fetch(ids []string) ([]byte, error) {
	res := []byte(`response for fetch task`)
	return res, nil
}

func (c Client) Delete(ids []string) error {
	return nil
}

func (c Client) Healthy(host string) bool {
	return true
}
