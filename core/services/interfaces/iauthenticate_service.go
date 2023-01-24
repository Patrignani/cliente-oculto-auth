package interfaces

import oauth "github.com/Patrignani/simple-oauth"

type IAuthenticateService interface {
	ClientCredentialsAuthorization(client *oauth.OAuthClient) oauth.AuthorizationRolesClient
	PasswordAuthorization(pass *oauth.OAuthPassword) oauth.AuthorizationRolesPassword
	RefreshTokenCredentialsAuthorization(refresh *oauth.OAuthRefreshToken) oauth.AuthorizationRolesRefresh
}
