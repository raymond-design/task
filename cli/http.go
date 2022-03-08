package cli

import "net/http"

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
