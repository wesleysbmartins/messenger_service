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
	log.Write("Message Received", message)

	err := usecases.ReceiveMessage(*message)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("Failed To Proccess Message!\nERROR: %s", err)
		json.NewEncoder(w).Encode(msg)
		log.Write(msg, message)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		msg := "Message Received With Success!"
		json.NewEncoder(w).Encode(msg)
		log.Write(msg, message)
	}
}
