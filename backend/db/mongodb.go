package db

import (
	"context"
	"fmt"
	"gobitly/configs"
	"gobitly/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
	err        error
)

const dbname = "gobitly"
const colname = "redirect"

func Connect() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(configs.EnvMongoDB()).SetServerAPIOptions(serverAPI)

	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	collection = client.Database(dbname).Collection(colname)

	// Send a ping to confirm a successful connection
	var result bson.M
	if err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func Disconnect() {
	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func InsertGobitly(gobitly models.Gobitly) {
	inserted, err := collection.InsertOne(context.Background(), gobitly)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 gobitly with id: ", inserted.InsertedID)
}

func UpdateGobitlyClick(gobitlyId string) {
	id, err := primitive.ObjectIDFromHex(gobitlyId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{"clicked": 1}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

func DeleteGobitly(gobitlyId string) {
	id, err := primitive.ObjectIDFromHex(gobitlyId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted count: ", result.DeletedCount)
}

func GetAllGobitlies() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var gobitlies []bson.M

	for cursor.Next(context.Background()) {
		var gobitly bson.M
		if err := cursor.Decode(&gobitly); err != nil {
			log.Fatal(err)
		}
		gobitlies = append(gobitlies, gobitly)
	}

	return gobitlies
}
