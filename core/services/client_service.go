package services

import (
	"github.com/Patrignani/cliente-oculto-auth/core/entity"
	repository "github.com/Patrignani/cliente-oculto-auth/core/repository/interfaces"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/specifications"
	"github.com/Patrignani/cliente-oculto-auth/core/services/interfaces"
)

type ClientService struct {
	repository repository.IClientRepository
}

func NewClientService(repository repository.IClientRepository) interfaces.IClientService {
	return &ClientService{repository: repository}
}

func (c *ClientService) Authenticate(clientId string, clientSecret string) (*entity.Client, error) {
	specification := specifications.NewFindClientByClientIdAndClientSecret(clientId, clientSecret, map[string]int{"_id": 1})
	return c.repository.FindBySpecification(specification)
}
