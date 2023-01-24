package interfaces

import "github.com/Patrignani/cliente-oculto-auth/core/entity"

type IRefreshTokenSerice interface {
	CreateRefreshToken(userID string) (*entity.RefreshToken, error)
	FindById(refreshTokenId string) (*entity.RefreshToken, error)
}
