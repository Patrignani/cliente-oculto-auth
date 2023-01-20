package main

import (
	"net/http"
	"strconv"

	common "github.com/Patrignani/cliente-oculto-auth/core/config"
	"github.com/Patrignani/cliente-oculto-auth/core/facades"
	oauth "github.com/Patrignani/simple-oauth"
	"github.com/labstack/echo/v4"
)

func main() {
	authFacade := facades.CreateFacade()

	expireTimeMinutesClient, err := strconv.Atoi(common.JwtExpireTimeMinutesClient)

	if err != nil {
		panic(err)
	}

	expireTimeMinutes, err := strconv.Atoi(common.JwtExpireTimeMinutes)

	if err != nil {
		panic(err)
	}

	options := &oauth.OAuthSimpleOption{
		Key:                     common.JwtKey,
		ExpireTimeMinutesClient: expireTimeMinutesClient,
		ExpireTimeMinutes:       expireTimeMinutes,
		AuthRouter:              common.AuthRouter,
	}

	authConfigure := &oauth.OAuthConfigure{
		ClientCredentialsAuthorization: authFacade.AuthenticateService.ClientCredentialsAuthorization,
	}

	e := echo.New()

	authRouter := oauth.NewAuthorization(authConfigure, options, e)

	authRouter.CreateAuthRouter()
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "I aaaaaaaaaaaaaaaaaaaaaaaAm Foda")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
