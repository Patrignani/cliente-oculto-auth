package interfaces

import (
	"github.com/Patrignani/cliente-oculto-auth/core/entity"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/specifications"
)

type IUserRepository interface {
	FindOneBySpecification(specification specifications.ISpecificationByOne) (*entity.User, error)
}
