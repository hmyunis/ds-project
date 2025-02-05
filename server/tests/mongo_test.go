package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//mogodb connection test

func TestMongoDBConnection(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://root:password1234@cluster0.21sbi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	assert.Nil(t, err)

	collection := client.Database("chatapp").Collection("messages")
	_, err = collection.InsertOne(context.TODO(), bson.M{"sender": "user1", "message": "Hello"})
	assert.Nil(t, err)
}
