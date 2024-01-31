package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	// "github.com/farmforum/controllers"
	"github.com/farmforum/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HandleAddUser(db *mongo.Database) http.Handler {
	userCollection := db.Collection("users")
	var userDocument models.User
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {

		err := json.NewDecoder(r.Body).Decode(&userDocument)
		if err != nil {
			panic(err)
		}

		currentTime := time.Now()
		userDocument.CreatedAt = currentTime
		userDocument.UpdatedAt = currentTime
		hashedPassword, err := HashPassword(userDocument.Password)
		if err != nil {
			panic(err)
		}
		userDocument.Password = hashedPassword

		_, err = userCollection.InsertOne(context.TODO(), userDocument)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusCreated)

	}))
}

