package facades

import (
	"context"
	"log"
	"time"

	common "github.com/Patrignani/cliente-oculto-auth/core/config"
	"github.com/Patrignani/cliente-oculto-auth/core/data"
	"github.com/Patrignani/cliente-oculto-auth/core/repository"
	"github.com/Patrignani/cliente-oculto-auth/core/services"
	"github.com/Patrignani/cliente-oculto-auth/core/services/interfaces"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthFacade struct {
	ClientService       interfaces.IClientService
	UserService         interfaces.IUserService
	RefreshTokenService interfaces.IRefreshTokenSerice
	AuthenticateService interfaces.IAuthenticateService
}

func CreateFacade() *AuthFacade {
	mongoContext := getMongoContext()

	//repo
	clientRepository := repository.NewClientRepository(mongoContext)
	userRepository := repository.NewUserRepository(mongoContext)
	refreshTokenRepository := repository.NewRefreshTokenRepository(mongoContext)

	//services
	clientService := services.NewClientService(clientRepository)
	userService := services.NewUserService(userRepository)
	refreshTokenService := services.NewRefreshTokenService(refreshTokenRepository)
	authServices := services.NewAuthenticateService(clientService, userService, refreshTokenService)

	return &AuthFacade{clientService, userService, refreshTokenService, authServices}
}

func getMongoContext() data.IMongoContext {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mongo := data.GetInstance()
	credential := options.Credential{
		Username:      common.MongodbUser,
		Password:      common.MongodbPassword,
		PasswordSet:   true,
		AuthSource:    common.MongodbDatabase,
		AuthMechanism: common.MongodbAuth,
	}

	if err := mongo.Initialize(ctx, credential, "mongodb://"+common.MongodbHosts+":"+common.MongodbPort,
		common.MongodbDatabase, &common.MongodbReplicaset); err != nil {
		log.Println("Could not resolve Data access layer", err)
	}

	return mongo
}
