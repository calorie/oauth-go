package domain

import (
	"errors"
)

type Client struct {
	ClientId     string
	ClientName   string
	RedirectUris []string
}

func (c *Client) FindRedirectUri(uri string) (string, error) {
	for _, u := range c.RedirectUris {
		if u == uri {
			return u, nil
		}
	}
	return "", errors.New("redirect_uri is wrong")
}
