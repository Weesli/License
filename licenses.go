package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type License struct {
	Owner  string `bson:"owner"`
	Name   string `bson:"name"`
	Key    string `bson:"key"`
	Status string `bson:"status"`
}

func getLicenses() []License {
	client := getConnection()
	collection := client.Database("test").Collection("licenses")

	var licenses []License
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var license License
		err := cursor.Decode(&license)
		if err != nil {
			panic(err)
		}
		licenses = append(licenses, license)
	}

	if err := cursor.Err(); err != nil {
		panic(err)
	}

	return licenses
}
func getLicense(key string) *License {
	client := getConnection()
	collection := client.Database("test").Collection("licenses")

	var license License
	filter := bson.D{{Key: "key", Value: key}}

	err := collection.FindOne(context.Background(), filter).Decode(&license)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}

	return &license
}

func createLicense(license License) {
	client := getConnection()
	collection := client.Database("test").Collection("licenses")

	_, err := collection.InsertOne(context.Background(), license)
	if err != nil {
		panic(err)
	}
	fmt.Println("License created successfully")
}
