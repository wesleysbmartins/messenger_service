package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	Id       primitive.ObjectID `bson:"_id"`
	Received MessageReceived    `json:"received" bson:"received"`
	Send     MessageSend        `json:"send" bson:"send"`
	Status   Status             `json:"status" bson:"status"`
}

type Status string

const (
	Received Status = "Received"
	Sended   Status = "Sended"
	Errored  Status = "Errored"
)
