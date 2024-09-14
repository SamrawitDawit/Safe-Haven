package repositories

import (
	"backend/domain"
	"backend/usecases/interfaces"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database, collectionName string) interfaces.UserRepositoryInterface {
	return &UserRepository{
		collection: db.Collection(collectionName),
	}
}

func (u *UserRepository) CreateUser(user *domain.User) error {
	_, err := u.collection.InsertOne(context.TODO(), user)
	return err
}

func (u *UserRepository) UpdateUser(user *domain.User) error {
	user.UpdatedAt = time.Now()
	fmt.Println(user.ID)
	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": bson.M{
			"accessToken":  user.AccessToken,
			"refreshToken": user.RefreshToken,
			"updatedAt":    user.UpdatedAt,
		},
	}
	updatedCount, err := u.collection.UpdateOne(context.TODO(), filter, update)
	fmt.Println(updatedCount)
	return err
}

func (u *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"email": email}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (u *UserRepository) GetUserByPhoneNumber(phoneNumber string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"phoneNumber": phoneNumber}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (u *UserRepository) GetUserByAnonymousDifferentiator(differentiator string) (*domain.User, error) {
	var anonUser domain.User
	filter := bson.M{"anonymousDifferentiator": differentiator}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&anonUser)
	if err != nil {
		return nil, err
	}
	return &anonUser, nil
}

func (u *UserRepository) GetUserByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"_id": id}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return &user, err
}

func (u *UserRepository) GetUsersCount() (int, error) {
	count, err := u.collection.CountDocuments(context.TODO(), bson.M{})
	return int(count), err
}
