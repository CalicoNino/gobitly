package db

import (
	"context"
	"fmt"
	"gobitly/configs"
	"gobitly/models"

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
	if err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func Disconnect() {
	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func InsertGobitly(gobitly models.Gobitly) (*mongo.InsertOneResult, error) {
	inserted, err := collection.InsertOne(context.Background(), gobitly)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func UpdateGobitlyClick(gobitlyId string) (*mongo.UpdateResult, error) {
	id, err := primitive.ObjectIDFromHex(gobitlyId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{"clicked": 1}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteGobitly(gobitlyId string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(gobitlyId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetGobitly(gobitlyId string) (*models.Gobitly, error) {
	id, err := primitive.ObjectIDFromHex(gobitlyId)
	var gobitly *models.Gobitly
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}

	err = collection.FindOne(context.Background(), filter).Decode(&gobitly)
	if err != nil {
		return nil, err
	}

	return gobitly, nil
}

func GetAllGobitlies() ([]bson.M, error) {
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var gobitlies []bson.M

	for cursor.Next(context.Background()) {
		var gobitly bson.M
		if err := cursor.Decode(&gobitly); err != nil {
			return nil, err
		}
		gobitlies = append(gobitlies, gobitly)
	}

	return gobitlies, nil
}
