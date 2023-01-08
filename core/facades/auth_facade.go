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
	AuthenticateService interfaces.IAuthenticateService
}

func CreateFacade() *AuthFacade {
	mongoContext := getMongoContext()

	//repo
	clientRepository := repository.NewClientRepository(mongoContext)

	//services
	clientService := services.NewClientService(clientRepository)
	authServices := services.NewAuthenticateService(clientService)

	return NewAuthFacade(clientService, authServices)
}

func NewAuthFacade(client interfaces.IClientService, authenticate interfaces.IAuthenticateService) *AuthFacade {
	return &AuthFacade{client, authenticate}
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
