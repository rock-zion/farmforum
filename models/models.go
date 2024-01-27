package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Crop struct {
		Id              primitive.ObjectID `bson:"_id,omitempty"`
		Name            string             `bson:"name,omitempty"`
		GestationInDays int16              `bson:"gestationInDays,omitempty"`
		BotanicalName   string             `bson:"botanicalName,omitempty"`
		Description     string             `bson:"description,omitempty"`
	}

	Question struct {
		Id            primitive.ObjectID `bson:"_id,omitempty"`
		Question      string             `bson:"question"`
		CropId        primitive.ObjectID `bson:"cropId,omitempty"`
		OlderVersions []string           `bson:"olderVersions,omitempty"`
		CreatedAt     time.Time          `bson:"createdAt"`
		UpdatedAt     time.Time          `bson:"updatedAt"`
	}

	Answer struct {
		Id         primitive.ObjectID `bson:"_id,omitempty"`
		QuestionId primitive.ObjectID `bson:"question"`
		Reply      string             `bson:"reply"`
		CreatedAt  time.Time          `bson:"createdAt"`
	}
)
