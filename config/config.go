package config

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var MI MongoInstance

var Database struct {
	URI      string
	Database string
	Timeout  time.Duration
}

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}
