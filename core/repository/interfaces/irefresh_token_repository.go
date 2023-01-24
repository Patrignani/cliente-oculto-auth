package interfaces

import (
	"github.com/Patrignani/cliente-oculto-auth/core/entity"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/specifications"
)

type IRefreshToken interface {
	Insert(refreshToken *entity.RefreshToken) error
	FindOneBySpecification(specification specifications.ISpecificationByOne) (*entity.RefreshToken, error)
	Update(filter map[string]interface{}, fields interface{}) error
}
