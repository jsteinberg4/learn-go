package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jsteinberg4/learn-go/mongodb-usage-examples/datatypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOneExample(client *mongo.Client) {
	coll := client.Database("sample_restaurants").Collection("restaurants")
	filter := bson.D{{"name", "Bagels N Buns"}}

	var result datatypes.Restaurant
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query didn't match any docs
			// NOTE: could do "&& != mongo.ErrNoDocuments"
			return
		}
		panic(err)
	}

	// Print whole doc
	// output, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", output)

	// Proving that it's actually marshalled to that datatype
	fmt.Println("Restaurant(")
	fmt.Printf("\tName=%s\n\tRestaurantId=%s\n\tCuisine=%s\n)", result.Name, result.RestaurantId, result.Cuisine)
}

func FindManyExample(client *mongo.Client) {
	var (
		coll    *mongo.Collection
		filter  bson.D
		results []datatypes.Restaurant
		cursor  *mongo.Cursor
		err     error
	)
	coll = client.Database("sample_restaurants").Collection("restaurants")
	filter = bson.D{{"cuisine", "Italian"}}
	cursor, err = coll.Find(context.Background(), filter, options.Find().SetLimit(5))

	// Read the whole cursor into memory
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// Print v1: Decode
	for _, result := range results {
		fmt.Printf("Pre-decode: %T", result)
		fmt.Printf("Name: %s\n", result.Name)
		cursor.Decode(&result)
		fmt.Printf("Post-decode: %T", result)
		fmt.Printf("%v\n", result)
		break
	}
}

func main() {
	// Load environment vars
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	// Connect to mongo
	uri := os.Getenv(("MONGO_URI"))
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Run through examples
	// FindOneExample(client)
	FindManyExample(client)
}
