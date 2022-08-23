package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Database struct {
	client *mongo.Client
}

func (d *Database) Connect() {
	uri := os.Getenv("MONGODB")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	d.client = client
}
func (d *Database) Disconnect() {
	err := d.client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
