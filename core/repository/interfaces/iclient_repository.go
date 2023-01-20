package interfaces

import (
	"github.com/Patrignani/cliente-oculto-auth/core/entity"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/specifications"
)

type IClientRepository interface {
	Insert(client entity.Client) error
	FindOneBySpecification(specification specifications.ISpecification) (*entity.Client, error)
}
