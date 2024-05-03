package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/entity"
	"github.com/tesis/internal/rooms/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type roomRepository struct {
	dbInstance *mongo.Database
}

var collection *mongo.Collection

type RoomRepository interface {
	InsertOne(roomEntity *entity.Room) (string, error)
	FindAll() ([]entity.RoomRes, error)
	FindOne(id string) (*entity.RoomRes, error)
	UpdateOne(id string, roomEntity *entity.Room) (*entity.Room, error)
	DeleteOne(id string) (bool, error)

	AddUserToRoom(idRoom, idUser string) error
	RemoveUserInRoom(idRoom, idUser string) error
}

func NewRoomRepository(dbMongoInstance *mongo.Database) RoomRepository {
	collection = dbMongoInstance.Collection("rooms")
	return &roomRepository{
		dbInstance: dbMongoInstance,
	}
}

func (rr *roomRepository) InsertOne(roomEntity *entity.Room) (string, error) {
	roomModel := models.Room{}
	roomModel.MapEntityToModel(roomEntity)

	if !utils.CheckRoomStatus(roomModel.Status) {
		return "", fmt.Errorf("room status is no valid")
	}

	ctx := context.TODO()
	result, err := collection.InsertOne(ctx, roomModel)
	if err != nil {
		panic(err)
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("el ID insertado no es un ObjectID")
	}

	return insertedID.Hex(), nil
}

func (rr *roomRepository) FindAll() ([]entity.RoomRes, error) {
	var rooms []models.RoomRes
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &rooms)
	if err != nil {
		return nil, err
	}

	roomsEntity := []entity.RoomRes{}
	for i := 0; i < len(rooms); i++ {
		roomentity := rooms[i].MapEntityFromModel()
		roomsEntity = append(roomsEntity, *roomentity)
	}

	return roomsEntity, nil
}

func (rr *roomRepository) FindOne(id string) (*entity.RoomRes, error) {
	var room models.RoomRes

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&room)
	if err != nil {
		return nil, err
	}

	return room.MapEntityFromModel(), nil
}

func (rr *roomRepository) UpdateOne(id string, roomEntity *entity.Room) (*entity.Room, error) {
	room := models.Room{}
	room.MapEntityToModel(roomEntity)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	update := bson.M{
		"$set": room,
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("the document to update was not found")
	}

	roomEntity.ID = id

	return roomEntity, nil
}

func (rr *roomRepository) DeleteOne(id string) (bool, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		log.Println("Error deleting document:", err)
		return false, err
	}
	if result.DeletedCount == 0 {
		return false, nil
	}

	return true, nil
}

func (rr *roomRepository) AddUserToRoom(roomID, userID string) error {
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

func (rr *roomRepository) RemoveUserInRoom(roomID, userID string) error {
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
