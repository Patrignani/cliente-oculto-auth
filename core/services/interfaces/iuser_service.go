package interfaces

import "github.com/Patrignani/cliente-oculto-auth/core/entity"

type IUserService interface {
	Authenticate(username string, password string) (*entity.User, error)
	FindById(userId string) (*entity.User, error)
}
