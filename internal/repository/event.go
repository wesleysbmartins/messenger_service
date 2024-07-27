package repository

import (
	"context"
	mongo "messenger_service/internal/adapters/database/mongo/operations"
	"messenger_service/internal/entities"
	"messenger_service/internal/shared/exceptions"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventRepository struct{}

var exception exceptions.DatabaseException

type EventParams struct {
	Id     primitive.ObjectID `bson:"_id"`
	Status string             `bson:"Status"`
}

const eventCollection string = "Messages"

type IEventRepository interface {
	Create(event entities.Event) error
	Find(filter *EventParams) (*[]entities.Event, error)
	Update(id primitive.ObjectID, value entities.Event) error
	Delete(id primitive.ObjectID) error
	handleFilter(params *EventParams) bson.M
	handleSet(params *entities.Event) bson.M
}

func (r *EventRepository) Create(event entities.Event) error {
	database := &mongo.MongoOperations{}
	ctx := context.TODO()

	event.Id = primitive.NewObjectID()

	_, err := database.Insert(ctx, eventCollection, event)

	if err != nil {
		exception.Handle(ctx, "CREATE", nil, event, err)
	}

	return err
}

func (r *EventRepository) Find(params *EventParams) (*[]entities.Event, error) {
	database := &mongo.MongoOperations{}
	ctx := context.TODO()

	values := []entities.Event{}

	filter := r.handleFilter(params)

	err := database.Find(ctx, eventCollection, filter, &values)

	if err != nil {
		exception.Handle(ctx, "FIND", filter, nil, err)
	}

	return &values, err
}

func (r *EventRepository) Update(id primitive.ObjectID, value entities.Event) error {
	database := &mongo.MongoOperations{}
	ctx := context.TODO()

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	set := r.handleSet(&value)

	_, err := database.Update(ctx, eventCollection, filter, set)

	if err != nil {
		exception.Handle(ctx, "UPDATE", filter, set, err)
	}

	return err
}

func (r *EventRepository) Delete(id primitive.ObjectID) error {
	database := &mongo.MongoOperations{}
	ctx := context.TODO()

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	_, err := database.Delete(ctx, eventCollection, filter)

	if err != nil {
		exception.Handle(ctx, "DELETE", filter, nil, err)
	}

	return err
}

func (r *EventRepository) handleFilter(params *EventParams) bson.M {
	filter := bson.M{}

	if params != nil {

		if params.Id != primitive.NilObjectID {
			filter["_id"] = params.Id
		}

		if params.Status != "" {
			filter["Status"] = params.Status
		}
	}

	return filter
}

func (r *EventRepository) handleSet(params *entities.Event) bson.M {
	update := bson.M{}

	if params != nil {

		set := bson.M{}

		if params.Status != "" {
			set["Status"] = params.Status
		}

		if len(set) > 0 {
			update["$set"] = set
		}
	}

	return update
}
