package mongo

import (
	"context"
	"fmt"
	database "messenger_service/internal/services/database/mongo"
	"messenger_service/internal/shared/exceptions"
)

type MongoOperations struct{}

type IMongoOperations interface {
	Insert(ctx context.Context, value any) any
	Find(ctx context.Context, filter any)
	Update(ctx context.Context, filter any, update any) any
	Delete(ctx context.Context, filter any) any
}

func (m *MongoOperations) Insert(ctx context.Context, value any) error {
	_, err := database.Mongo.Collection.InsertOne(ctx, value)
	if err != nil {
		fmt.Println("ERR", err.Error())
		panic(err)
	}
	return err
}

func (m *MongoOperations) Find(ctx context.Context, filter interface{}, resultType interface{}) {
	cur, err := database.Mongo.Collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	cur.All(ctx, resultType)
}

func (m *MongoOperations) Update(ctx context.Context, filter interface{}, value interface{}) any {
	exception := &exceptions.DatabaseException{}
	result, err := database.Mongo.Collection.UpdateOne(ctx, filter, value)
	if err != nil {
		exception.Handle(ctx, "Update", filter, value, err)
	}
	return result
}

func (m *MongoOperations) Delete(ctx context.Context, filter any) any {
	result, err := database.Mongo.Collection.InsertOne(ctx, filter)
	if err != nil {
		panic(err)
	}
	return result
}
