package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/farmforum/config"
	"github.com/farmforum/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// this file houses all crud logic for questions

// edit a crop entry
func FetchCropById() http.Handler {
	cropCollection := db.DB().Collection("crop")
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			panic(err)
		}

		filter := bson.M{"_id": id}
		cursor, err := cropCollection.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}

		result := []models.Crop{}
		err = cursor.All(context.TODO(), &result)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-type", "application/vnd.api+json")
		json.NewEncoder(w).Encode(result)

	}))
}

// fetch all crops
func FetchAllCrops() http.Handler {
	cropCollection := db.DB().Collection("crop")
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
func AddCrop() http.Handler {
	cropDocument := models.Crop{}
	cropCollection := db.DB().Collection("crop")
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
func DeleteCrop() http.Handler {
	cropCollection := db.DB().Collection("crop")
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			panic(err)
		}
		result, err := cropCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
		if err != nil {
			panic(fmt.Errorf("this is the error 1 %v", err))
		}

		w.WriteHeader(http.StatusNoContent)
		log.Println(result)
	}))
}

// edit a crop entry
func EditCrop() http.Handler {
	cropCollection := db.DB().Collection("crop")
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Updating...")
		vars := mux.Vars(r)
		var cropToUpdate models.Crop
		err := json.NewDecoder(r.Body).Decode(&cropToUpdate)
		if err != nil {
			panic(err)
		}

		update := bson.M{
			"$set": bson.M{
				"gestationInDays": cropToUpdate.GestationInDays,
				"description":     cropToUpdate.Description},
		}

		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			panic(err)
		}
		singleResult, err := cropCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
		if err != nil {
			panic(err)
		}
		fmt.Println(singleResult)
	}))
}

// insert many a crop entry
func InsertManyCrops() http.Handler {
	var cropDocument = []interface{}{models.Crop{}}
	cropCollection := db.DB().Collection("crop")
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		println("Adding Many...")
		err := json.NewDecoder(r.Body).Decode(&cropDocument)
		if err != nil {
			panic(fmt.Errorf("this is the error 1 %v", err))
		}

		insertResult, err := cropCollection.InsertMany(context.TODO(), cropDocument)
		if err != nil {
			panic(fmt.Errorf("this is the error 2 %v", err))
		}
		w.WriteHeader(http.StatusCreated)
		log.Println(insertResult)

	}))
}
