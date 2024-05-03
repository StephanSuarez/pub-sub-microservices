package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type roomSubRepository struct {
	dbInstance *mongo.Database
}

var collection *mongo.Collection

type RoomSubRepository interface {
	AddUserToRoom(idRoom, idUser string) error
	RemoveUserInRoom(idRoom, idUser string) error
}

func NewRoomSubRepository(dbMongoInstance *mongo.Database) RoomSubRepository {
	collection = dbMongoInstance.Collection("rooms")
	return &roomSubRepository{
		dbInstance: dbMongoInstance,
	}
}

func (rr *roomSubRepository) AddUserToRoom(roomID, userID string) error {
	objectID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		log.Println(err)
		return err
	}

	filtro := bson.D{{Key: "_id", Value: objectID}}

	actualizacion := bson.D{
		{Key: "$addToSet", Value: bson.D{
			{Key: "users", Value: userID},
		}},
	}

	result, err := collection.UpdateOne(context.Background(), filtro, actualizacion)
	if err != nil {
		log.Fatal(err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("room not found or user already added")
	}

	return nil
}

func (rr *roomSubRepository) RemoveUserInRoom(roomID, userID string) error {
	log.Println(roomID)
	log.Println(userID)
	objectID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		log.Println(err)
		return err
	}

	filtro := bson.D{{Key: "_id", Value: objectID}}

	actualizacion := bson.D{
		{Key: "$pull", Value: bson.D{
			{Key: "users", Value: userID},
		}},
	}

	result, err := collection.UpdateOne(context.Background(), filtro, actualizacion)
	if err != nil {
		log.Fatal(err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("room not found or user already removed")
	}

	return nil
}
