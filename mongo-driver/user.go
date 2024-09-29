package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const collectionUsers = "users"

type User struct {
	Id       bson.ObjectID `bson:"_id"`
	Forename string        `bson:"forename"`
	Surname  string        `bson:"surname"`
	Age      int           `bson:"age"`
}

func buildRandomUsers(count int64) []User {
	random := rand.New(rand.NewSource(time.Now().UnixMilli()))
	users := make([]User, 0, count)
	for i := 0; i < int(count); i++ {
		users = append(users, User{
			Id:       bson.NewObjectID(),
			Forename: fmt.Sprintf("Forename %d", i),
			Surname:  fmt.Sprintf("Surname %d", i),
			Age:      random.Intn(80),
		})
	}
	return users
}

func createUsers(ctx context.Context, db *mongo.Database, users []User) error {
	col := db.Collection(collectionUsers)
	_, err := col.InsertMany(ctx, users)
	return err
}

func countUsers(ctx context.Context, db *mongo.Database) (int64, error) {
	col := db.Collection(collectionUsers)
	return col.CountDocuments(ctx, bson.D{})
}

func getUser(ctx context.Context, db *mongo.Database, filter interface{},
	opts ...options.Lister[options.FindOneOptions]) (*User, error) {
	col := db.Collection(collectionUsers)
	var user User
	err := col.FindOne(ctx, filter, opts...).Decode(&user)
	return &user, err
}

func updateUsers(ctx context.Context, db *mongo.Database, filter interface{},
	update interface{},
	opts ...options.Lister[options.UpdateOptions]) (int64, error) {
	col := db.Collection(collectionUsers)
	result, err := col.UpdateMany(ctx, filter, update, opts...)
	if err != nil {
		return 0, err
	} else {
		return result.MatchedCount, nil
	}
}

func deleteUser(ctx context.Context, db *mongo.Database, id bson.ObjectID) error {
	col := db.Collection(collectionUsers)
	_, err := col.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	return err
}

func deleteUsers(ctx context.Context, db *mongo.Database, filter interface{},
	opts ...options.Lister[options.DeleteOptions]) (int64, error) {
	col := db.Collection(collectionUsers)
	result, err := col.DeleteMany(ctx, filter, opts...)
	if err != nil {
		return 0, err
	} else {
		return result.DeletedCount, nil
	}
}
