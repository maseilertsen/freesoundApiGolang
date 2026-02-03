package handlers

import (
	"freesoundApiGolang/utils"
	"log"
	"net/http"
	"os"
	"strings"
)

func HandleSomething(w http.ResponseWriter, r *http.Request) {

	// Load API_key from .env
	api_key := os.Getenv("FREESOUND_API_KEY")

	// Get ID as a string
	id := strings.TrimPrefix(r.URL.Path, utils.SOUND_PATH)

	if id == "" {
		http.Error(w, "Please provide an ID", http.StatusBadRequest)
		return
	}

	getRequest := utils.API_V2 + id + "&token=" + api_key
	log.Println("Request URL with api: " + getRequest) // TODO remove debug

	// Get request towards FreesoundAPI with ID
	res, err := http.Get(getRequest)
	if err != nil {
		http.Error(w, "An unexpected error orccurred.", http.StatusInternalServerError)
		log.Println("Failed to fetch data with GET")
	}
	log.Println("From API: ") // todo remove
	log.Println(res)          // todo remove
	defer res.Body.Close()

}
