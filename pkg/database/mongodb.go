package database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"net/url"
	"os"
)

//const uri = "mongodb+srv://lavumi:<password>@cluster0.kuovpbb.mongodb.net/?retryWrites=true&w=majority"

type Database struct{}

func (db *Database) Init() {

	cluster := os.Getenv("CLUSTER")
	username := os.Getenv("USER")
	password := os.Getenv("PASS")

	uri := "mongodb+srv://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster +
		"/?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {

		panic(err)

	}

	//defer client.Disconnect(context.Background())

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}
