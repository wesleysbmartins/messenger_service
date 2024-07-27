package usecases

import (
	"messenger_service/internal/entities"
	"messenger_service/internal/mock"
	"messenger_service/internal/repository"
)

func ReceiveMessage(message entities.MessageReceived) error {
	repository := &repository.EventRepository{}

	event := &entities.Event{}
	event.Received = message
	event.Status = entities.Received

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

	err := repository.Create(*event)

	return err
}
