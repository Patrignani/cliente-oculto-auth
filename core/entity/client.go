package entity

import "time"

type Client struct {
	ID             string    `bson:"_id" json:"id"`
	ClientID       string    `bson:"client_id" json:"client_id"`
	ClientSecret   string    `bson:"client_secret" json:"client_secret"`
	Name           string    `bson:"name" json:"name"`
	Description    string    `bson:"description" json:"description"`
	CreateAt       time.Time `bson:"create_at" json:"create_at"`
	UpdateAt       time.Time `bson:"update_at" json:"update_at"`
	CreateBy       string    `bson:"create_by" json:"create_by"`
	ClientCreateBy string    `bson:"client_create_by" json:"client_create_by"`
	UpdateBy       string    `bson:"update_by" json:"update_by"`
	ClientUpdateBy string    `bson:"client_update_by" json:"client_update_by"`
	Active         bool      `bson:"active" json:"active"`
}
