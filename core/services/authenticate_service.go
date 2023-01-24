package services

import (
	"github.com/Patrignani/cliente-oculto-auth/core/services/interfaces"
	oauth "github.com/Patrignani/simple-oauth"
)

type AuthenticateService struct {
	clientService      interfaces.IClientService
	userService        interfaces.IUserService
	refreshTokenSerice interfaces.IRefreshTokenSerice
}

func NewAuthenticateService(client interfaces.IClientService, user interfaces.IUserService, refresh interfaces.IRefreshTokenSerice) interfaces.IAuthenticateService {
	return &AuthenticateService{clientService: client, userService: user, refreshTokenSerice: refresh}
}

func (a *AuthenticateService) ClientCredentialsAuthorization(client *oauth.OAuthClient) oauth.AuthorizationRolesClient {
	authorization := oauth.AuthorizationRolesClient{}

	clientAuth, err := a.clientService.Authenticate(client.Client_id, client.Client_secret)

	if err != nil {
		authorization.Authorized = false
	}

	authorization.Authorized = clientAuth != nil && len(clientAuth.ID) > 0

	return authorization
}

func (a *AuthenticateService) PasswordAuthorization(pass *oauth.OAuthPassword) oauth.AuthorizationRolesPassword {
	authorization := oauth.AuthorizationRolesPassword{}

	clientAuth, err := a.clientService.Authenticate(pass.Client_id, pass.Client_secret)

	if err != nil || clientAuth == nil {
		authorization.Authorized = false
	}

	user, err := a.userService.Authenticate(pass.Username, pass.Password)

	if err != nil || user == nil {
		authorization.Authorized = false
	} else {

		refresh, err := a.refreshTokenSerice.CreateRefreshToken(user.ID)

		if err != nil {
			authorization.Authorized = false
		} else {

			authorization.Authorized = true
			authorization.Subject = user.ID
			authorization.Permissions = user.Permissions
			authorization.Roles = user.Roles
			authorization.RefreshToken = refresh.ID
		}
	}

	return authorization
}

func (a *AuthenticateService) RefreshTokenCredentialsAuthorization(refresh *oauth.OAuthRefreshToken) oauth.AuthorizationRolesRefresh {
	authorization := oauth.AuthorizationRolesRefresh{}

	clientAuth, err := a.clientService.Authenticate(refresh.Client_id, refresh.Client_secret)

	if err != nil || clientAuth == nil {
		authorization.Authorized = false
	}

	refreshToken, err := a.refreshTokenSerice.FindById(refresh.Refresh_token)

	if refreshToken == nil || err != nil {
		authorization.Authorized = false
	} else {

		user, err := a.userService.FindById(refreshToken.UserID)

		if user == nil || err != nil {
			authorization.Authorized = false
		} else {
			refresh, err := a.refreshTokenSerice.CreateRefreshToken(user.ID)

			if err != nil {
				authorization.Authorized = false
			} else {

				authorization.Authorized = true
				authorization.Subject = user.ID
				authorization.Permissions = user.Permissions
				authorization.Roles = user.Roles
				authorization.RefreshToken = refresh.ID
			}
		}
	}
	return authorization
}
