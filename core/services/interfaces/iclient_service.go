package interfaces

import "github.com/Patrignani/cliente-oculto-auth/core/entity"

type IClientService interface {
	Authenticate(clientId string, clientSecret string) (*entity.Client, error)
}
