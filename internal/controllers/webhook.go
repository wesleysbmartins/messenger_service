package controllers

import (
	"encoding/json"
	"fmt"
	"messenger_service/internal/entities"
	"messenger_service/internal/shared/logger"
	usecases "messenger_service/internal/usecases/receive"
	"net/http"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	message := &entities.MessageReceived{}
	json.NewDecoder(r.Body).Decode(&message)

	log := &logger.Logger{}
	log.Read("Message Received", message)

	err := usecases.ReceiveMessage(*message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, http.StatusText(200))
}
