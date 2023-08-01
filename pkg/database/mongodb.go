package database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"grpc-game-server/pkg/model"
	"log"
	"net/url"
	"os"
)

//const uri = "mongodb+srv://lavumi:<password>@cluster0.kuovpbb.mongodb.net/?retryWrites=true&w=majority"

type MongoDb struct{}

var client *mongo.Client

func (db *MongoDb) Init() *mongo.Client {

	cluster := os.Getenv("CLUSTER")
	username := os.Getenv("USER")
	password := os.Getenv("PASS")

	uri := "mongodb+srv://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster +
		"/?retryWrites=true&w=majority"

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	//defer func() {
	//	if err = client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	filter := bson.D{{"name", "lavumi2"}}

	loginCollection := client.Database("user").Collection("login")

	var user model.Login
	err = loginCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User Query test %s | %s\n", user.Name, user.PlatformId)
	return client
}
