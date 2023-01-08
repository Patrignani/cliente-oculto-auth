package repository

import (
	"github.com/Patrignani/cliente-oculto-auth/core/data"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/interfaces"
)

type UserRepository struct {
	context data.IMongoContext
}

func NewUserRepository(context data.IMongoContext) interfaces.IUserRepository {
	return &UserRepository{context: context}
}
