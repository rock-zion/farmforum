package main

import (
	"context"
	"log"
	"net/http"
	"os"


	"github.com/farmforum/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa 1024
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey, signKey []byte
)

func init() {
	var err error
	signKey, err = os.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	verifyKey, err = os.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := mux.NewRouter()
	router.Handle("/crops", controllers.AddCrop(famrmforumDatabase)).Methods("POST")
	router.Handle("/crops", controllers.FetchAllCrops(famrmforumDatabase)).Methods("GET")
	router.Handle("/crops/{id}", controllers.FetchCropById(famrmforumDatabase)).Methods("GET")
	router.Handle("/crops/many", controllers.InsertManyCrops(famrmforumDatabase)).Methods("POST")
	router.Handle("/crops/{id}", controllers.DeleteCrop(famrmforumDatabase)).Methods("DELETE")
	router.Handle("/crops/{id}", controllers.EditCrop(famrmforumDatabase)).Methods("PUT")
	router.Handle("/crops/questions/{id}", controllers.FetchQuestionByCropId(famrmforumDatabase)).Methods("GET")

	//
	router.Handle("/question/{id}", controllers.AddQuestion(famrmforumDatabase)).Methods("POST")
	router.Handle("/questions/{id}", controllers.UpdateQuestion(famrmforumDatabase)).Methods("GET")

	// routers to handle user actions
	router.Handle("/user", controllers.HandleAddUser(famrmforumDatabase)).Methods(("POST"))
	router.Handle("/user/login", controllers.HandleLogUserIn(famrmforumDatabase)).Methods(("POST"))

	server := &http.Server{
		Addr:    os.Getenv("SERVER_ADDR"),
		Handler: router,
	}

	server.ListenAndServe()
}
