package handlers

import (
	"encoding/json"
	"fmt"
	"freesoundApiGolang/models"
	"freesoundApiGolang/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func HandleSongId(w http.ResponseWriter, r *http.Request) {
	// Set CORS header
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)

	// Load API_key from .env
	api_key := os.Getenv("FREESOUND_API_KEY")

	// Get ID as a string
	providedId := strings.TrimPrefix(r.URL.Path, utils.SOUND_PATH)
	id := strings.TrimSuffix(providedId, "/")
	// ID is empty
	if id == "" {
		http.Error(w, "Please provide an ID", http.StatusBadRequest)
		return
	}
	// ID is not numeric
	log.Println("ID provided: " + id)
	if stringIsInt(id) != 1 {
		http.Error(w, "ID's can only be numbers.", http.StatusBadRequest)
		log.Println("ID provided is not numeric")
		return
	}

	getRequest := utils.API_V2 + id + "&token=" + api_key + utils.MINIMAL_FIELDS
	//log.Println("Request URL with api: " + getRequest) // TODO remove debug

	// Get request towards FreesoundAPI with ID
	res, err := http.Get(getRequest)
	if err != nil {
		http.Error(w, "An unexpected error orccurred.", http.StatusInternalServerError)
		log.Println("Failed to fetch data with GET")
		return
	}
	defer res.Body.Close()

	var songInfo models.SongInfo
	if err := json.NewDecoder(res.Body).Decode(&songInfo); err != nil {
		http.Error(w, "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		log.Println("Failed to Marshal json (HandleSongId")
		return
	}

	// Check if there is any results
	if len(songInfo.Results) == 0 {
		http.Error(w, "No results found for this ID", http.StatusNotFound)
		return
	}

	// Build response
	response := models.SingleSong{
		Artist: songInfo.Results[0].Username,
		Song:   songInfo.Results[0].Name,
	}

	//log.Printf("artist: %v \nName of song: %v\n", response.Artist, response.Song)// TODO remove debug

	// Send JSON response to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HandleMultipleSongs fetches multiple songs by comma-separated IDs
// Usage: /sounds?ids=123,456,789
func HandleMultipleSongs(w http.ResponseWriter, r *http.Request) {
	// Set CORS header
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)

	api_key := os.Getenv("FREESOUND_API_KEY")

	// Get IDs from query parameter
	idsParam := r.URL.Query().Get("ids")
	if idsParam == "" {
		http.Error(w, "Please provide IDs using ?ids=123,456,789", http.StatusBadRequest)
		return
	}

	// Split into individual IDs
	ids := strings.Split(idsParam, ",")
	var results models.MultipleSongs

	// Fetch each song sequentially
	for _, id := range ids {
		id = strings.TrimSpace(id)

		// Validate ID is numeric
		if stringIsInt(id) != 1 {
			log.Printf("Skipping invalid ID: %s", id)
			continue
		}

		// Fetch from Freesound API
		getRequest := utils.API_V2 + id + "&token=" + api_key + utils.MINIMAL_FIELDS
		res, err := http.Get(getRequest)
		if err != nil {
			log.Printf("Failed to fetch ID %s: %v", id, err)
			continue
		}

		var songInfo models.SongInfo
		if err := json.NewDecoder(res.Body).Decode(&songInfo); err != nil {
			log.Printf("Failed to decode ID %s: %v", id, err)
			res.Body.Close()
			continue
		}
		res.Body.Close()

		// Extract and store result
		if len(songInfo.Results) > 0 {
			results = append(results, models.SingleSong{
				Artist: songInfo.Results[0].Username,
				Song:   songInfo.Results[0].Name,
			})
		}
	}

	// Output as plain text
	w.Header().Set("Content-Type", "text/plain")

	//fmt.Fprintf(w, "<Artist> - <Song>\n") //TODO: remove debug
	for _, song := range results {
		fmt.Fprintf(w, "%s - %s\n", song.Artist, song.Song)
	}
}

// Returns 1 if string is int
// Return -1 if string isnt int
func stringIsInt(id string) int {
	if _, err := strconv.Atoi(id); err != nil {
		return -1
	}
	return 1
}
