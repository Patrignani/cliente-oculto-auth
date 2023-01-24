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
	userCollection = "users"
)

type UserRepository struct {
	context data.IMongoContext
}

func NewUserRepository(context data.IMongoContext) interfaces.IUserRepository {
	return &UserRepository{context: context}
}

func (u *UserRepository) FindOneBySpecification(specification specifications.ISpecificationByOne) (*entity.User, error) {
	filter, opts := specification.GetSpecification()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var user entity.User
	if mgoErr := u.context.FindOne(ctx, userCollection, filter, &user, opts); mgoErr != nil {
		return nil, mgoErr
	}

	return &user, nil
}
