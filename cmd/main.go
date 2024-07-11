package main

import (
	gateways "messenger_service/internal/gateways/meta"
	"messenger_service/internal/mock"
	database "messenger_service/internal/services/database/mongo"
	"messenger_service/internal/services/dotenv"
	"messenger_service/internal/services/http"
	"messenger_service/internal/services/server"
	"messenger_service/internal/shared/logger"
	usecases "messenger_service/internal/usecases/send"
)

func init() {
	dotenv := &dotenv.DotEnv{}
	dotenv.Load()

	log := &logger.Logger{}
	log.Load()

	mongo := &database.MongoDB{}
	mongo.Connect()

	mockSingleton := &mock.Mock{}
	mockSingleton.Load()

	client := &http.Client{}
	client.Load()

	metaGateway := &gateways.MetaGateway{}
	metaGateway.Load()
}

func main() {
	go usecases.SendMessage()

	server := &server.Server{}
	server.Run()
}
