package handlers

import (
	"fmt"
	"freesoundApiGolang/utils"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	// Offer information for redirection to paths
	output := "This service does not provide any functionality on root path level. Please use " +
		"<a href=\"" + utils.SOUNDS_PATH + "\">" + utils.SOUNDS_PATH + "?ids=123,456" + "</a>"

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)
	log.Println("Entered root handler " + utils.ROOT_PATH)

	if err != nil {
		http.Error(w, "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
	}
}
