package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Login struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	PlatformId string             `bson:"platform_id"`
}
