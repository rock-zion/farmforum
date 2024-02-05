package main

import (
	"log"
	"net/http"
	"os"

	"github.com/farmforum/routes"
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

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	server := &http.Server{
		Addr:    os.Getenv("SERVER_ADDR"),
		Handler: routes.Routes(),
	}

	server.ListenAndServe()
}
