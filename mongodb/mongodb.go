package mongodb

import (
	"context"
	"fmt"

	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// Connect connect to mongo server
func Connect(host, user, password, dbName, mechanism, source string) error {
	connectOptions := options.ClientOptions{}
	// Set auth if have
	if user != "" && password != "" {
		connectOptions.Auth = &options.Credential{
			AuthMechanism: mechanism,
			AuthSource:    source,
			Username:      user,
			Password:      password,
		}
	}

	// Connect
	client, err := mongo.Connect(context.Background(), connectOptions.ApplyURI(host))
	if err != nil {
		fmt.Println("Error when connect to MongoDB database", host, err)
		return err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO MONGODB: " + host))

	// Set data
	db = client.Database(dbName)
	return nil
}

// GetInstance ...
func GetInstance() *mongo.Database {
	return db
}