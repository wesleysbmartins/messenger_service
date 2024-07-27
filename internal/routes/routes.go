package routes

import (
	"encoding/json"
	"fmt"
	"messenger_service/internal/controllers"
	"messenger_service/internal/entities"
	"messenger_service/internal/shared/logger"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Dracarys!")
}

func meta(w http.ResponseWriter, r *http.Request) {

	type value struct {
		url     *url.URL
		message entities.MessageSend
	}

	message := &entities.MessageSend{}
	json.NewDecoder(r.Body).Decode(&message)

	url := r.URL

	valueToLog := &value{
		url:     url,
		message: *message,
	}

	log := &logger.Logger{}
	log.Write("Meta Received", valueToLog)

	fmt.Fprintln(w, "Meta Received")
}

func Routes(router *mux.Router) {

	router.HandleFunc("/", healthCheck).Methods("GET")

	router.HandleFunc("/meta", meta).Methods("POST")

	router.HandleFunc("/webhook", controllers.Webhook).Methods("POST")
}
