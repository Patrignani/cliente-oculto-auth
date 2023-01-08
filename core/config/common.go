package config

import (
	"os"
)

var (
	MongodbUser       = os.Getenv("MONGODB_USER")
	MongodbPassword   = os.Getenv("MONGODB_PASSWORD")
	MongodbHosts      = os.Getenv("MONGODB_HOST")
	MongodbPort       = os.Getenv("MONGODB_PORT")
	MongodbAuth       = os.Getenv("MONGODB_AUTH")
	MongodbDatabase   = os.Getenv("MONGODB")
	MongodbReplicaset = os.Getenv("MONGODB_REPLICASET")
)
