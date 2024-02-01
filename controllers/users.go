package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	// "github.com/farmforum/controllers"
	"github.com/farmforum/models"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
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

	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		userCollection := db.Collection("users")
		var userDocument models.User

		err := json.NewDecoder(r.Body).Decode(&userDocument)
		if err != nil {
			panic(err)
		}

		singleResult := userCollection.FindOne(context.TODO(), bson.M{"email": strings.ToLower(userDocument.Email)})
		if singleResult != nil {
			w.WriteHeader(http.StatusConflict)
			return
		}

		currentTime := time.Now()
		userDocument.CreatedAt = currentTime
		userDocument.UpdatedAt = currentTime
		userDocument.Email = strings.ToLower(userDocument.Email)
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

func HandleLogUserIn(db *mongo.Database) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		userCollection := db.Collection("users")
		usersDocument := models.User{}
		filterResponseDoc := models.User{}

		err := json.NewDecoder(r.Body).Decode(&usersDocument)
		if err != nil {
			panic(err)
		}

		err = userCollection.FindOne(context.TODO(), bson.M{"email": usersDocument.Email}).Decode(&filterResponseDoc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			if CheckPasswordHash(usersDocument.Password, filterResponseDoc.Password) {
				key, _ := os.ReadFile("keys/app.rsa")
				signKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
				if err != nil {
					panic(err)
				}
				// create a signer for rsa 256
				token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
					"email": filterResponseDoc.Email,
					"name":  filterResponseDoc.Name,
				})

				signedString, err := token.SignedString(signKey)
				if err != nil {
					panic(err)
				}
				log.Println(signedString, err)

				w.WriteHeader(http.StatusAccepted)

				res := map[string]interface{}{
					"data": map[string]interface{}{},
				}

				res["token"] = signedString
				res["data"].(map[string]interface{})["name"] = filterResponseDoc.Name
				res["data"].(map[string]interface{})["id"] = filterResponseDoc.Id
				res["data"].(map[string]interface{})["email"] = filterResponseDoc.Email
				res["data"].(map[string]interface{})["createdAt"] = filterResponseDoc.
					CreatedAt
				res["data"].(map[string]interface{})["updatedAt"] = filterResponseDoc.UpdatedAt

				json.NewEncoder(w).Encode(res)
				return
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
	}))
}
