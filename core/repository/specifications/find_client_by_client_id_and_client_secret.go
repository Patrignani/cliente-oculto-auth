package specifications

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type FindClientByClientIdAndClientSecret struct {
	clientId     string
	clientSecret string
	project      map[string]int
}

func NewFindClientByClientIdAndClientSecret(clientId string, clientSecret string, project map[string]int) ISpecificationByOne {
	return &FindClientByClientIdAndClientSecret{clientId, clientSecret, project}
}

func (t *FindClientByClientIdAndClientSecret) GetSpecification() (map[string]interface{}, *options.FindOneOptions) {
	opts := options.FindOne().
		SetProjection(t.project)

	filter := bson.M{"client_id": t.clientId, "client_secret": t.clientSecret, "active": true}

	return filter, opts
}
