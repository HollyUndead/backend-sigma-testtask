package database

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	. "testtry2/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var userCollection *mongo.Collection
var COLLECTION = "user"

func GetClient() *mongo.Client {
	uri := os.Getenv("DATABASE_URL")
	if client != nil {
		return client
	}
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	if userCollection != nil {
		return userCollection
	}
	userCollection := client.Database("golangtestproj").Collection(collectionName)
	return userCollection
}

func Disconnect() {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	if client == nil {
		return
	}
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func List_Users() []User {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	var userList []User
	cursor, err := userCollection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for cursor.Next(ctx) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		userList = append(userList, user)
	}
	return userList
}

func Find_User_ByFirstName(firstName string) *User {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	var user *User
	filter := bson.D{{Key: "FirstName", Value: firstName}}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

func Find_User_ById(id primitive.ObjectID) *User {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	var user *User
	filter := bson.D{{Key: "_id", Value: id}}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

func Delete_User(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	filter := bson.M{"_id": id}
	result, err := userCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	if result.DeletedCount == 0 {
		return result, nil
	}
	return result, err
}

func Create_User(user User) *User {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	newId := primitive.NewObjectID()
	userToPost := User{
		Id:          newId,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Age:         user.Age,
		PhoneNumber: user.PhoneNumber,
	}
	result, err := userCollection.InsertOne(ctx, userToPost)
	if err != nil {
		return nil
	}
	userNew := Find_User_ById(newId)
	fmt.Println(result)
	return userNew
}

func Update_User(user User) (any, int) {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	filter := bson.M{"_id": user.Id}
	update := bson.D{{Key: "$set",
		Value: bson.D{{Key: "firstname", Value: user.FirstName},
			{Key: "lastname", Value: user.LastName},
			{Key: "age", Value: user.Age},
			{Key: "phonenumber", Value: user.PhoneNumber}}}}
	result, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	if result.MatchedCount == 0 {
		return "User does not exist", http.StatusAccepted
	}
	updatedUser := Find_User_ById(user.Id)
	return updatedUser, http.StatusOK
}
