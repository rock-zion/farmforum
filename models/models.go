package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Crop struct {
		Id              primitive.ObjectID `bson:"_id,omitempty"`
		Name            string             `bson:"name,omitempty"`
		GestationInDays int16              `bson:"gestationInDays,omitempty"`
		BotanicalName   string             `bson:"botanicalName,omitempty"`
		Description     string             `bson:"description,omitempty"`
	}
)
