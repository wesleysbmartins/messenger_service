package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	host     string
	port     string
	user     string
	password string
	database string
}

type IMongoDB interface {
	Connect()
}

var Database *mongo.Database

func (m *MongoDB) Connect() {
	if Database == nil {
		credentials := &MongoDB{
			user:     os.Getenv("MONGO_USER"),
			password: os.Getenv("MONGO_PASSWORD"),
			database: os.Getenv("MONGO_DATABASE"),
			port:     os.Getenv("MONGO_PORT"),
			host:     os.Getenv("MONGO_HOST"),
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

		Database = client.Database(credentials.database)

		fmt.Println("DATABASE CONNECTION SUCCESS!")
	}
}
