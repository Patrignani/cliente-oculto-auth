package entity

type Client struct {
	Id           string `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Description  string `json:"description,omitempty" bson:"description,omitempty"`
	ClientId     string `json:"client_id,omitempty" bson:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty" bson:"client_secret,omitempty"`
	Active       bool   `json:"active" bson:"active"`
}
