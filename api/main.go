package main

import (
	"net/http"
	"strconv"
	"strings"

	common "github.com/Patrignani/cliente-oculto-auth/core/config"
	"github.com/Patrignani/cliente-oculto-auth/core/facades"
	oauth "github.com/Patrignani/simple-oauth"
	t "github.com/golang-jwt/jwt/v5"
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
		CustomActionRolesMiddleware:          customActionRolesMiddleware,
	}

	e := echo.New()

	authRouter := oauth.NewAuthorization(authConfigure, options, e)

	authRouter.CreateAuthRouter()

	jwtValidate := authRouter.GetDefaultMiddleWareJwtValidate()

	e.GET("/health", func(c echo.Context) error {
		// Execute verificações de saúde aqui
		// Por exemplo, verificar a conexão com o banco de dados, serviços externos, etc.

		// Se tudo estiver bem, responda com um status OK (200)
		return c.String(http.StatusOK, "Health check passed")
	})

	g := e.Group("check-jwt")
	g.Use(jwtValidate)
	g.GET("", func(c echo.Context) error {
		id := c.Get("user-id")
		cid := c.Get("cid")
		get := c.Get("user")

		println(cid, id)

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

		ID := claims["user-id"].(string)
		return c.String(http.StatusOK, "Id:"+ID+" roles:"+strings.Join(rolesStr, ",")+" permissions:"+strings.Join(permissionsStr, ","))
	}, authRouter.RolesMiddleware("10", "5"))

	e.Logger.Fatal(e.Start(":8000"))
}

func customActionRolesMiddleware(c echo.Context, token *t.Token, claims t.MapClaims) error {
	if claims["user-id"] != nil {
		c.Set("user-id", claims["user-id"].(string))
	}

	if claims["cid"] != nil {
		c.Set("cid", claims["cid"].(string))
	}

	return nil
}
