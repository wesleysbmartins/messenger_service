package usecases

import (
	"messenger_service/internal/entities"
	gateways "messenger_service/internal/gateways/meta"
	"messenger_service/internal/repository"
	"messenger_service/internal/shared/logger"
	"os"
	"strconv"
	"time"
)

func SendMessage() {
	log := &logger.Logger{}
	gateway := &gateways.MetaGateway{}
	repo := &repository.EventRepository{}
	timeToSleep, _ := strconv.Atoi((os.Getenv("TIME_TO_SLEEP")))

	eventParams := &repository.EventParams{
		Status: string(entities.Received),
	}

	for {

		events, err := repo.Find(eventParams)
		log.Write("FIND EVENTS ERROR", err)

		for _, event := range *events {

			eventUpdate := entities.Event{}

			err = gateway.SendMessage(event.Send)

			if err != nil {
				eventUpdate = entities.Event{Status: entities.Sended}
				log.Write("SEND MESSAGE ERROR", err)
			} else {
				eventUpdate = entities.Event{Status: entities.Errored}
				log.Write("SEND MESSAGE SUCCESS", err)
			}

			err = repo.Update(event.Id, eventUpdate)
			log.Write("UPDATE EVENT ERROR", err)
		}

		time.Sleep(time.Duration(timeToSleep) * time.Minute)
	}
}
