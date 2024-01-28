package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/farmforum/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// this file houses all crud logic for questions
func AddQuestion(farmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		questionDocument := models.Question{}
		questionCollection := farmforumDatabase.Collection("questions")

		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			panic(err)
		}

		err = json.NewDecoder(r.Body).Decode(&questionDocument)
		if err != nil {
			panic(err)
		}
		currentTime := time.Now()
		questionDocument.CreatedAt = currentTime
		questionDocument.UpdatedAt = currentTime
		questionDocument.CropId = id

		_, err = questionCollection.InsertOne(context.TODO(), questionDocument)
		w.WriteHeader(http.StatusCreated)

	}))
}

func UpdateQuestion(farmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}

func FetchQuestionById() http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}
