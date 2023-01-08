package main

import (
	"github.com/Patrignani/cliente-oculto-auth/core/facades"
	oauth "github.com/Patrignani/simple-oauth"
	"github.com/labstack/echo/v4"
)

func main() {
	authFacade := facades.CreateFacade()

	options := &oauth.OAuthSimpleOption{
		Key:               "teste",
		ExpireTimeMinutes: 10,
		AuthRouter:        "/Auth",
	}

	authConfigure := &oauth.OAuthConfigure{
		ClientCredentialsAuthorization: authFacade.AuthenticateService.ClientCredentialsAuthorization,
	}

	e := echo.New()

	authRouter := oauth.NewAuthorization(authConfigure, options, e)

	authRouter.CreateAuthRouter()

	e.Logger.Fatal(e.Start(":8000"))
}