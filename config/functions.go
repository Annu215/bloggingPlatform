package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Init() {
	Database.URI = "mongodb://localhost:27017"
	Database.Database = "blog"
	Database.Timeout = 10 * time.Second

}
func ConnectDB() {
	opts := options.Client()
	opts.ApplyURI(Database.URI)
	opts.ConnectTimeout = &Database.Timeout
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalln("Error in connecting to mongo URI :", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), Database.Timeout)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln("Error in connecting:", err)
	}
	log.Println("Database connected..!!")
	MI = MongoInstance{
		Client: client,
		DB:     client.Database(Database.Database),
	}
}
