package repository

import (
	"errors"

	"github.com/calorie/oauth-go/domain"
)

type ClientRepository struct {
}

func NewClientRepositoty() *ClientRepository {
	return &ClientRepository{}
}

func (r *ClientRepository) FindClient(cid string) (*domain.Client, error) {
	for _, c := range *r.clients() {
		if c.ClientId == cid {
			return &c, nil
		}
	}
	return nil, errors.New("client_id is wrong")
}

func (r *ClientRepository) clients() *[]domain.Client {
	return &[]domain.Client{
		{
			ClientId: "1",
			ClientName: "client1",
			RedirectUris: []string{"https://example.com"},
		},
	}
}
