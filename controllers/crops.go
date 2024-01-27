package controllers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// this file houses all crud logic for questions

// fetch all crops
func FetchAllCrops(famrmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}

// add a crop entry
func AddCrop(famrmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}

// delete a crop entry
func DeleteCrop(famrmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}

// edit a crop entry
func EditCrop(famrmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}

// insert many a crop entry
func InsertManyCrops(famrmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}
