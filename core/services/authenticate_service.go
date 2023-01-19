package services

import (
	"github.com/Patrignani/cliente-oculto-auth/core/services/interfaces"
	oauth "github.com/Patrignani/simple-oauth"
)

type AuthenticateService struct {
	clientService interfaces.IClientService
}

func NewAuthenticateService(client interfaces.IClientService) interfaces.IAuthenticateService {
	return &AuthenticateService{clientService: client}
}

func (a *AuthenticateService) ClientCredentialsAuthorization(client *oauth.OAuthClient) oauth.AuthorizationRolesClient {
	Authorization := new(oauth.AuthorizationRolesClient)

	clientAuth, err := a.clientService.Authenticate(client.Client_id, client.Client_secret)

	if err != nil {
		Authorization.Authorized = false
	}

	Authorization.Authorized = clientAuth != nil && len(clientAuth.ID) > 0

	return *Authorization
}
