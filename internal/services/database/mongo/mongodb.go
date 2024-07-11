package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	port            string
	host            string
	user            string
	password        string
	collection_name string
	database_name   string
	Collection      *mongo.Collection
}

type IMongoDB interface {
	Connect()
}

var Mongo *MongoDB

func (m *MongoDB) Connect() {

	if Mongo == nil {
		credentials := &MongoDB{
			user:            os.Getenv("MONGO_USER"),
			password:        os.Getenv("MONGO_PASSWORD"),
			collection_name: os.Getenv("MONGO_COLLECTION"),
			database_name:   os.Getenv("MONGO_DATABASE"),
			port:            os.Getenv("MONGO_PORT"),
			host:            os.Getenv("MONGO_HOST"),
		}

		uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", credentials.user, credentials.password, credentials.host, credentials.port)

		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

		ctx := context.TODO()
		client, err := mongo.Connect(ctx, opts)

		if err != nil {
			fmt.Println("URI:", uri)
			panic(err)
		}

		pingErr := client.Ping(ctx, nil)

		if pingErr != nil {
			fmt.Println("PING ERROR:", pingErr.Error())
			panic(err)
		}

		credentials.Collection = client.Database(credentials.database_name).Collection(credentials.collection_name)
		Mongo = credentials

		fmt.Println("DATABASE CONNECTION SUCCESS!")
	}
}
