package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Crop struct {
		Id              primitive.ObjectID `bson:"_id,omitempty"`
		Name            string             `json:"name,omitempty"`
		GestationInDays int16              `json:"gestationInDays,omitempty"`
		BotanicalName   string             `json:"botanicalName,omitempty"`
		Description     string             `json:"description,omitempty"`
	}

	Question struct {
		Id            primitive.ObjectID `bson:"_id,omitempty"`
		Question      string             `json:"question"`
		CropId        primitive.ObjectID `bson:"cropId,omitempty"`
		OlderVersions []string           `json:"olderVersions,omitempty"`
		CreatedAt     time.Time          `json:"createdAt"`
		UpdatedAt     time.Time          `json:"updatedAt"`
	}

	Answer struct {
		Id         primitive.ObjectID `bson:"_id,omitempty"`
		QuestionId primitive.ObjectID `json:"question"`
		Reply      string             `json:"reply"`
		CreatedAt  time.Time          `json:"createdAt"`
	}
)
