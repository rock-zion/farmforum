package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/farmforum/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// this file houses all crud logic for questions

// fetch all crops
func FetchAllCrops(farmforumDatabase *mongo.Database) http.Handler {
	cropCollection := farmforumDatabase.Collection("crop")
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		cursor, err := cropCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			panic(err)
		}
		var crops []models.Crop
		if err = cursor.All(context.TODO(), &crops); err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/vnd.api+json")
		json.NewEncoder(w).Encode(crops)

	}))
}

// add a crop entry
func AddCrop(farmforumDatabase *mongo.Database) http.Handler {
	cropDocument := models.Crop{}
	cropCollection := farmforumDatabase.Collection("crop")
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&cropDocument)
		if err != nil {
			panic(err)
		}

		insertResult, err := cropCollection.InsertOne(context.TODO(), cropDocument)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusCreated)
		log.Println(insertResult)
	}))
}

// delete a crop entry
func DeleteCrop(farmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}

// edit a crop entry
func EditCrop(farmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}

// insert many a crop entry
func InsertManyCrops(farmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}
