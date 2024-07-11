package usecases

import (
	"context"
	"fmt"
	"messenger_service/internal/entities"
	gateways "messenger_service/internal/gateways/meta"
	mongo "messenger_service/internal/services/database/mongo/operations"
	"messenger_service/internal/shared/logger"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SendMessage() {
	log := &logger.Logger{}
	events := &[]entities.Event{}
	database := &mongo.MongoOperations{}
	gateway := &gateways.MetaGateway{}

	timeToSleep, _ := strconv.Atoi((os.Getenv("TIME_TO_SLEEP")))

	filter := bson.D{primitive.E{Key: "status", Value: entities.Received}}

	for {
		fmt.Println("START EXECUTION")
		ctx := context.TODO()
		database.Find(ctx, filter, events)

		for _, event := range *events {
			err := gateway.SendMessage(event.Send)

			filter := bson.M{"_id": event.Id}

			if err == nil {
				log.Read("MESSAGE SEND WITH SUCCESS!", event.Send)

				update := bson.M{"$set": bson.M{"status": entities.Sended}}

				database.Update(ctx, filter, update)
			} else {
				log.Read("ERROR TO SEND MESSAGE!", event.Send)

				update := bson.M{"$set": bson.M{"status": entities.Errored}}

				database.Update(ctx, filter, update)
			}
		}

		time.Sleep(time.Duration(timeToSleep) * time.Minute)
	}
}
