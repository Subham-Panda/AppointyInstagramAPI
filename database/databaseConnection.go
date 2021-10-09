package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database is a wrapper around the Mongo Client
type DatabaseConnection struct {
	Client 	*mongo.Client
	Ctx 	context.Context
}

// ConnectDatabase connects to the database and returns a DatabaseConnection
func ConnectDatabase() *DatabaseConnection {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db := DatabaseConnection{Client: mongoClient, Ctx: mongoCtx}
	err = db.Client.Connect(db.Ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected To Database")
	return &db
}

// DisconnectDatabase disconnects from the database
func DisconnectDatabase(db *DatabaseConnection) {
	db.Client.Disconnect(db.Ctx)
	fmt.Println("Disconnected From Database")
}