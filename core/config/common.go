package config

import (
	"os"
)

var (
	AuthRouter                 = os.Getenv("AUTH_ROUTER")
	JwtExpireTimeMinutesClient = os.Getenv("JWT_EXPIRE_CLIENT")
	JwtExpireTimeMinutes       = os.Getenv("JWT_EXPIRE")
	JwtKey                     = os.Getenv("JWT_KEY")
	MongodbUser                = os.Getenv("MONGODB_USER")
	MongodbPassword            = os.Getenv("MONGODB_PASSWORD")
	MongodbHosts               = os.Getenv("MONGODB_HOST")
	MongodbPort                = os.Getenv("MONGODB_PORT")
	MongodbAuth                = os.Getenv("MONGODB_AUTH")
	MongodbDatabase            = os.Getenv("MONGODB")
	MongodbReplicaset          = os.Getenv("MONGODB_REPLICASET")
)
