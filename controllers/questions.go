package controllers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// this file houses all crud logic for questions

func AddQuestion(farmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}

func UpdateQuestion(farmforumDatabase *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}

func FetchQuestionById() http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {}))
}
