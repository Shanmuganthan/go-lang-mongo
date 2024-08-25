package common

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "connection_string"

const dbName = "go-lang-mongo"

var dbConnectionString *mongo.Database

var dbClinet *mongo.Client

func init() {

	fmt.Println("Connection Started")

	clientOption := options.Client().ApplyURI(connectionString)

	dbClinet, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal("Database connection Error", err)
	}

	log.Print("Database Connected")

	dbConnectionString = dbClinet.Database(dbName)

}

func GetDB() *mongo.Database {
	return dbConnectionString
}

func GetClient() *mongo.Client {
	return dbClinet
}
