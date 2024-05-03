package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/entity"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/repository/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userSubRepository struct {
	dbInstance *mongo.Database
}

var collection *mongo.Collection

type UserSubRepository interface {
	GetUserByUserName(userName string) (*entity.UsersRes, error)
	GetUserByEmail(email string) (*entity.UsersRes, error)

	AddRoomToList(userID, roomID string) error
	RemoveRoomInList(userID, roomID string) error
}

func NewUserSubRepository(dbInstance *mongo.Database) UserSubRepository {
	collection = dbInstance.Collection("users")
	return &userSubRepository{
		dbInstance: dbInstance,
	}
}

func (rr *userSubRepository) GetUserByUserName(userName string) (*entity.UsersRes, error) {
	var userModel models.UsersR
	log.Println(userName)

	err := collection.FindOne(context.TODO(), bson.M{"userName": userName}).Decode(&userModel)
	if err != nil {
		return nil, err
	}

	return userModel.MapEntityFromModel(), nil
}

func (rr *userSubRepository) GetUserByEmail(email string) (*entity.UsersRes, error) {
	var userModel models.UsersR
	log.Println(email)
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&userModel)
	if err != nil {
		return nil, err
	}

	return userModel.MapEntityFromModel(), nil
}

func (rr *userSubRepository) AddRoomToList(userID, roomID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	update := bson.D{
		{Key: "$addToSet", Value: bson.D{
			{Key: "rooms", Value: roomID},
		}},
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("room not found or user already added")
	}

	return nil
}

func (rr *userSubRepository) RemoveRoomInList(userID, roomID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	update := bson.D{
		{Key: "$pull", Value: bson.D{
			{Key: "rooms", Value: roomID},
		}},
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("room not found or user already added")
	}

	return nil
}
