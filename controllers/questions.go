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
	"gopkg.in/mgo.v2/bson"
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
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		questionDocument := models.Question{}
		questionCollection := farmforumDatabase.Collection("questions")
		vars := mux.Vars(r)

		err := json.NewDecoder(r.Body).Decode(&questionDocument)
		if err != nil {
			panic(err)
		}
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			panic(err)
		}

		// Find an element vith the ID from request
		var result models.Question
		err = questionCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
		if err != nil {
			panic(err)
		}

		update := bson.M{
			"$set":  bson.M{"question": questionDocument.Question},
			"$push": bson.M{"olderVersions": result.Question},
		}

		_, err = questionCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
		if err != nil {
			panic(err)
		}

	}))
}

func FetchQuestionByCropId(farmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		questionCollection := farmforumDatabase.Collection("questions")

		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			panic(err)
		}
		filter := bson.M{"cropId": id}

		cursor, err := questionCollection.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}

		var questions []models.Question
		if err = cursor.All(context.TODO(), &questions); err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/vnd.api+json")
		json.NewEncoder(w).Encode(questions)

	}))
}
