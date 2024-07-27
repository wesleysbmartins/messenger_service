package main

import (
	database "messenger_service/internal/adapters/database/mongo"
	"messenger_service/internal/adapters/server"
	gateways "messenger_service/internal/gateways/meta"
	"messenger_service/internal/mock"
	"messenger_service/internal/shared/logger"
	usecases "messenger_service/internal/usecases/send"
	"messenger_service/pkg/dotenv"
	"messenger_service/pkg/http"
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
