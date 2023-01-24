package repository

import (
	"context"
	"time"

	"github.com/Patrignani/cliente-oculto-auth/core/data"
	"github.com/Patrignani/cliente-oculto-auth/core/entity"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/interfaces"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/specifications"
)

const (
	ClientCollection = "clients"
)

type ClientRepository struct {
	context data.IMongoContext
}

func NewClientRepository(context data.IMongoContext) interfaces.IClientRepository {
	return &ClientRepository{context: context}
}

func (c *ClientRepository) Insert(client *entity.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	id, err := c.context.Insert(ctx, ClientCollection, client)
	if err != nil {
		return err
	}

	client.ID = id

	return nil
}

func (u *ClientRepository) FindOneBySpecification(specification specifications.ISpecificationByOne) (*entity.Client, error) {
	filter, opts := specification.GetSpecification()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var client entity.Client
	if mgoErr := u.context.FindOne(ctx, ClientCollection, filter, &client, opts); mgoErr != nil {
		return nil, mgoErr
	}

	return &client, nil
}
