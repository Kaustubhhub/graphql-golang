package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client
}

var connectionString string = "mongodb://localhost:27017"

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil
	}

	if err := client.Ping(ctx, nil); err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
		return nil
	}

	return &DB{Client: client}
}
