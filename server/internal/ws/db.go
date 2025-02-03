package ws

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Database

func ConnectDB(mongoURI string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database("chatapphub")
	log.Println("Connected to MongoDB")
}

func GetRooms() (map[string]*Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Collection("rooms")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	rooms := make(map[string]*Room)

	for cursor.Next(ctx) {
		var room Room
		if err := cursor.Decode(&room); err != nil {
			log.Println(err)
			continue
		}
		rooms[room.ID] = &room
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func SaveRoom(room *Room) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.Collection("rooms")
	filter := bson.M{"id": room.ID}

	update := bson.M{"$set": room}

	upsert := true
	opts := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := collection.UpdateOne(ctx, filter, update, &opts)
	return err
}

func DeleteRoom(roomID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Collection("rooms")
	_, err := collection.DeleteOne(ctx, bson.M{"id": roomID})
	return err
}
