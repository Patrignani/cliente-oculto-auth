package interfaces

import oauth "github.com/Patrignani/simple-oauth"

type IAuthenticateService interface {
	ClientCredentialsAuthorization(client *oauth.OAuthClient) oauth.AuthorizationRolesClient
}
