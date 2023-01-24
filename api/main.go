package main

import (
	"net/http"
	"strconv"
	"strings"

	common "github.com/Patrignani/cliente-oculto-auth/core/config"
	"github.com/Patrignani/cliente-oculto-auth/core/facades"
	oauth "github.com/Patrignani/simple-oauth"
	t "github.com/golang-jwt/jwt"
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
		ClientCredentialsAuthorization:       authFacade.AuthenticateService.ClientCredentialsAuthorization,
		PasswordAuthorization:                authFacade.AuthenticateService.PasswordAuthorization,
		RefreshTokenCredentialsAuthorization: authFacade.AuthenticateService.RefreshTokenCredentialsAuthorization,
	}

	e := echo.New()

	authRouter := oauth.NewAuthorization(authConfigure, options, e)

	authRouter.CreateAuthRouter()
	jwtValidate := authRouter.GetDefaultMiddleWareJwtValidate()

	g := e.Group("CheckJwt")
	g.Use(jwtValidate)
	g.GET("", func(c echo.Context) error {
		get := c.Get("user")
		user := get.(*t.Token)
		claims := user.Claims.(t.MapClaims)
		roles := claims["roles"].([]interface{})

		//permissions := claims["permissions"].([]interface{})
		rolesStr := []string{}
		permissionsStr := []string{}
		for _, role := range roles {
			rolesStr = append(rolesStr, role.(string))
		}

		// for _, permission := range permissions {
		// 	permissionsStr = append(permissionsStr, permission.(string))
		// }

		ID := claims["sub"].(string)
		return c.String(http.StatusOK, "Id:"+ID+" roles:"+strings.Join(rolesStr, ",")+" permissions:"+strings.Join(permissionsStr, ","))
	}, authRouter.PermissionAndRoleMiddleware("1,2", "5,6,1"))

	e.Logger.Fatal(e.Start(":8000"))
}
