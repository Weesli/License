package controller

import (
	"LicenseChecker/database"
	"LicenseChecker/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Weesli")
	}
}

func GetLicenseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := r.URL.Query().Get("key")
	client := database.GetConnection()
	collection := client.Database("test").Collection("licenses")

	var license *model.License
	filter := bson.D{{Key: "key", Value: key}}

	err := collection.FindOne(context.Background(), filter).Decode(&license)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}
	if license == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "License not found"}`)
		return
	}
	json.NewEncoder(w).Encode(license)
}

func CreateLicenseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("admin-secret") == "" || r.Header.Get("admin-secret") != os.Getenv("ADMIN_SECRET") {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "Unauthorized"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var license model.License
	err := json.NewDecoder(r.Body).Decode(&license)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid request body"}`)
		return
	}
	client := database.GetConnection()
	collection := client.Database("test").Collection("licenses")

	_, err = collection.InsertOne(context.Background(), license)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(license)
}

func GetLicensesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("admin-secret") == "" || r.Header.Get("admin-secret") != os.Getenv("ADMIN_SECRET") {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "Unauthorized"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	client := database.GetConnection()
	collection := client.Database("test").Collection("licenses")

	var licenses []model.License
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var license model.License
		err := cursor.Decode(&license)
		if err != nil {
			panic(err)
		}
		licenses = append(licenses, license)
	}

	if err := cursor.Err(); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(licenses)
}

func DeleteLicenseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("admin-secret") == "" || r.Header.Get("admin-secret") != os.Getenv("ADMIN_SECRET") {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "Unauthorized"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	key := r.URL.Query().Get("key")
	client := database.GetConnection()
	collection := client.Database("test").Collection("licenses")

	filter := bson.D{{Key: "key", Value: key}}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents\n", result.DeletedCount)
	fmt.Fprintf(w, `{"message": "License deleted successfully"}`)
}
