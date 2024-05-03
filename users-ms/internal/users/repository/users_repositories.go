package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/entity"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/repository/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	dbInstance *mongo.Database
}

var collection *mongo.Collection

type UserRepository interface {
	InsertOne(userEntity *entity.Users) (string, error)
	FindAll() ([]entity.UsersRes, error)
	FindOne(id string) (*entity.UsersRes, error)
	UpdateOne(id string, userEntity *entity.Users) (*entity.UsersRes, error)
	DeleteOne(id string) (bool, error)
	GetUserByUserName(userName string) (*entity.UsersRes, error)
	GetUserByEmail(email string) (*entity.UsersRes, error)
}

func NewUserRepository(dbInstance *mongo.Database) UserRepository {
	collection = dbInstance.Collection("users")
	return &userRepository{
		dbInstance: dbInstance,
	}
}

func (ur *userRepository) InsertOne(userEntity *entity.Users) (string, error) {
	userModel := models.Users{}
	userModel.MapEntityToModel(userEntity)

	ctx := context.TODO()
	result, err := collection.InsertOne(ctx, userModel)
	if err != nil {
		panic(err)
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("el ID insertado no es un ObjectID")
	}

	return insertedID.Hex(), nil
}

func (ur *userRepository) FindAll() ([]entity.UsersRes, error) {
	users := []models.UsersR{}

	ctx := context.TODO()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	usersEntity := []entity.UsersRes{}
	for i := 0; i < len(users); i++ {
		userentity := users[i].MapEntityFromModel()
		usersEntity = append(usersEntity, *userentity)
	}

	return usersEntity, nil
}

func (rr *userRepository) FindOne(id string) (*entity.UsersRes, error) {
	var user models.UsersR

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user.MapEntityFromModel(), nil
}

func (rr *userRepository) UpdateOne(id string, userEntity *entity.Users) (*entity.UsersRes, error) {
	user := models.Users{}
	user.MapEntityToModel(userEntity)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	update := bson.M{
		"$set": user,
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("can not update the document")
	}

	userEntityRes := entity.UsersRes{
		ID:       id,
		UserName: userEntity.UserName,
		Email:    userEntity.Email,
	}

	return &userEntityRes, nil
}

func (rr *userRepository) DeleteOne(id string) (bool, error) {
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

func (rr *userRepository) GetUserByUserName(userName string) (*entity.UsersRes, error) {
	var userModel models.UsersR
	log.Println(userName)

	err := collection.FindOne(context.TODO(), bson.M{"userName": userName}).Decode(&userModel)
	if err != nil {
		return nil, err
	}

	return userModel.MapEntityFromModel(), nil
}

func (rr *userRepository) GetUserByEmail(email string) (*entity.UsersRes, error) {
	var userModel models.UsersR
	log.Println(email)
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&userModel)
	if err != nil {
		return nil, err
	}

	return userModel.MapEntityFromModel(), nil
}
