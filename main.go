package main

import (
	"freesoundApiGolang/handlers"
	"freesoundApiGolang/utils"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load() // looks for root .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Port has not been set, using default 8080")
		port = "8080"
	}

	http.HandleFunc("GET "+utils.ROOT_PATH, handlers.RootHandler)
	http.HandleFunc("GET "+utils.SOUND_PATH+"{id}/", handlers.HandleSomething)

	log.Println("Starting server on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
