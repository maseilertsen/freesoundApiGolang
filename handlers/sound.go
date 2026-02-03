package handlers

import (
	"freesoundApiGolang/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func HandleSongId(w http.ResponseWriter, r *http.Request) {

	// Load API_key from .env
	api_key := os.Getenv("FREESOUND_API_KEY")

	// Get ID as a string
	id := strings.TrimPrefix(r.URL.Path, utils.SOUND_PATH)

	// ID is empty
	if id == "" {
		http.Error(w, "Please provide an ID", http.StatusBadRequest)
		return
	}
	// ID is not numeric
	if stringIsInt(id) != 1 {
		http.Error(w, "ID's can only be numbers.", http.StatusBadRequest)
		log.Println("ID provided is not numeric")
		return
	}

	getRequest := utils.API_V2 + id + "&token=" + api_key
	log.Println("Request URL with api: " + getRequest) // TODO remove debug

	// Get request towards FreesoundAPI with ID
	res, err := http.Get(getRequest)
	if err != nil {
		http.Error(w, "An unexpected error orccurred.", http.StatusInternalServerError)
		log.Println("Failed to fetch data with GET")
		return
	}
	defer res.Body.Close()

	// TODO: check that ID is only numeric
	// TODO: Check that only one item is returned

}

// Returns 1 if string is int
// Return -1 if string isnt int
func stringIsInt(id string) int {
	if _, err := strconv.Atoi(id); err != nil {
		return -1
	}
	return 1
}
