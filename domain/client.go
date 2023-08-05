package domain

import (
	"errors"
)

type Client struct {
	ClientId     string
	ClientName   string
	RedirectUris []string
}

func FindClient(cid string) (*Client, error) {
	for _, c := range *clients() {
		if c.ClientId == cid {
			return &c, nil
		}
	}
	return nil, errors.New("client_id is wrong")
}

func clients() *[]Client {
	return &[]Client{
		{
			ClientId: "1",
			ClientName: "client1",
			RedirectUris: []string{"https://example.com"},
		},
	}
}

func (c *Client) FindRedirectUri(uri string) (string, error) {
	for _, u := range c.RedirectUris {
		if u == uri {
			return u, nil
		}
	}
	return "", errors.New("redirect_uri is wrong")
}
