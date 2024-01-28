package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/farmforum/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongo_uri := os.Getenv("MONGO_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongo_uri).SetServerAPIOptions(serverAPI)

	// create new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)

	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	famrmforumDatabase := client.Database(os.Getenv("FARM_FORUM_DB_NAME"))

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

	server := &http.Server{
		Addr:    os.Getenv("SERVER_ADDR"),
		Handler: router,
	}

	server.ListenAndServe()
}
