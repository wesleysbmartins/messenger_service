package usecases

import (
	"context"
	"messenger_service/internal/entities"
	"messenger_service/internal/mock"
	mongo "messenger_service/internal/services/database/mongo/operations"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReceiveMessage(message entities.MessageReceived) error {
	database := &mongo.MongoOperations{}

	event := &entities.Event{}
	event.Received = message
	event.Status = entities.Received
	event.Id = primitive.NewObjectID()

	for i := 0; i < len(message.Messaging); i++ {
		msg := message.Messaging[i]

		if msg.Message.Text != "" {
			for j := 0; j < len(mock.MockSingleton.Messages); j++ {
				resp := mock.MockSingleton.Messages[j]

				if resp.Message.Text != "" {
					event.Send = resp
				}
			}
		} else if msg.Postback.Payload != "" {
			for j := 0; j < len(mock.MockSingleton.Messages); j++ {
				resp := mock.MockSingleton.Messages[j]

				if resp.Message.Attachment.Payload.Text != "" {
					event.Send = resp
				}
			}
		}
	}

	ctx := context.TODO()
	err := database.Insert(ctx, event)

	return err
}
