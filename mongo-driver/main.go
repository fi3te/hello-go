package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const uri = "mongodb://localhost:27017"
const username = "root"
const password = "root"
const connectTimeoutInSeconds = 2
const programTimeoutInSeconds = 5

const minimumUserCount = 500

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*programTimeoutInSeconds)
	defer cancel()

	client, err := connectToDb(connectTimeoutInSeconds)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	log.Printf("Connected to MongoDB with URI '%s'.\n", uri)

	db := client.Database("demo")
	createdCount, err := createUsersIfNecessary(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created %d users.\n", createdCount)

	user, err := readOneOfTheOldestUsers(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read one of the oldest users: %v\n", *user)

	err = deleteUser(ctx, db, user.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted previously selected user.")

	deletedCount, err := deleteUnderageUsers(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted %d underage users.\n", deletedCount)

	years := 2
	err = increaseAgeOfAllUsers(ctx, db, years)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Increased age of users by %d years.\n", years)
}

func connectToDb(timeoutInSeconds int) (*mongo.Client, error) {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      username,
		Password:      password,
	}
	connectTimeout := time.Second * time.Duration(timeoutInSeconds)
	return mongo.Connect(options.Client().ApplyURI(uri).SetAuth(credential).SetConnectTimeout(connectTimeout))
}

func createUsersIfNecessary(ctx context.Context, db *mongo.Database) (int64, error) {
	count, err := countUsers(ctx, db)
	if err != nil {
		return 0, err
	}
	numberToCreate := minimumUserCount - count
	if numberToCreate < 1 {
		return 0, nil
	}

	users := buildRandomUsers(numberToCreate)
	err = createUsers(ctx, db, users)
	if err != nil {
		return 0, err
	}
	return numberToCreate, nil
}

func readOneOfTheOldestUsers(ctx context.Context, db *mongo.Database) (*User, error) {
	return getUser(ctx, db, bson.D{}, options.FindOne().SetSort(bson.D{{Key: "age", Value: -1}}))
}

func deleteUnderageUsers(ctx context.Context, db *mongo.Database) (int64, error) {
	return deleteUsers(ctx, db, bson.D{{Key: "age", Value: bson.D{{Key: "$lt", Value: 18}}}})
}

func increaseAgeOfAllUsers(ctx context.Context, db *mongo.Database, years int) error {
	_, err := updateUsers(ctx, db, bson.D{}, bson.D{{Key: "$inc", Value: bson.D{{Key: "age", Value: years}}}})
	return err
}
