package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func getUsers(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	var (
		coll   *mongo.Collection
		cursor *mongo.Cursor
		err    error
		docs   bson.A
	)
	enableCors(&w)

	coll = db.Collection("crud-users")
	cursor, err = coll.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	cursor.All(context.Background(), &docs)
	data, err := json.Marshal(docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
func getUser(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	fmt.Fprintf(w, "Get a user")
	log.Default().Println("Get a user")
}
func createUser(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	enableCors(&w)
	user := User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	coll := db.Collection("crud-users")
	result, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		defer log.Fatal(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Default().Printf("Inserted document with _id: %v\n", result.InsertedID)
		fmt.Fprintf(w, "Inserted document with _id: %v\n", result.InsertedID)
	}
}
func updateUser(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	fmt.Fprintf(w, "Update a user")
	log.Default().Println("Update a user")
}
func deleteUser(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	fmt.Fprintf(w, "Delete a user")
	log.Default().Println("Delete a user")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

// makeMongoHandler Wraps an http.HandlerFunc in a closure to connect and disconnect from MongoDB
func makeMongoHandler(fn func(http.ResponseWriter, *http.Request, *mongo.Database)) http.HandlerFunc {
	URI, found := os.LookupEnv("MONGO_URL")
	if !found {
		panic("No value MONGO_URL in environment")
	}

	// Connect to mongo
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	var db *mongo.Database = client.Database("sandbox")

	return func(w http.ResponseWriter, r *http.Request) {
		// Disconnect from mongo when the handler returns
		// defer client.Disconnect(context.Background())
		fn(w, r, db)
	}
}

func corsTestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		enableCors(&w)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		w.Header().Set("Access-Control-Max-Age", "10000")
	case http.MethodGet:
		enableCors(&w)
		data1 := map[string]int{
			"id":    0,
			"name":  1,
			"email": 2,
		}
		data3 := map[string]int{
			"id":    1,
			"name":  1,
			"email": 2,
		}
		data2 := []map[string]int{data1, data3}
		data_encoding, err := json.Marshal(data2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write(data_encoding)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	default:
		http.Error(w, "Endpoint does not support method "+r.Method, http.StatusMethodNotAllowed)
	}
}

func main() {

	URI, found := os.LookupEnv("MONGO_URL")
	if !found {
		panic("No value MONGO_URL in environment")
	}

	// Connect to mongo
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	var db *mongo.Database = client.Database("sandbox")
	coll := db.Collection("crud-users")
	// docs := bson.A{}
	filter := bson.D{}
	cursor, err := coll.Find(context.Background(), filter)
	// cursor.All(context.Background(), &docs)
	cursor.Next(context.Background())
	doc := cursor.Current
	fmt.Printf("Mongo results: %v\n", doc)

	// http.HandleFunc("/cors-test", corsTestHandler)
	// http.HandleFunc("/users", makeMongoHandler(getUsers))
	// http.HandleFunc("/users/", makeMongoHandler(getUser))
	// http.HandleFunc("/users/create", makeMongoHandler(createUser))
	// http.HandleFunc("/users/update", makeMongoHandler(updateUser))
	// http.HandleFunc("/users/delete", makeMongoHandler(deleteUser))
	//
	// http.ListenAndServe(":8080", nil)
}
