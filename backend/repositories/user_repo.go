package repositories

import (
	"backend/domain"
	"backend/usecases/interfaces"
	"context"
	"time"

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
	filter := bson.M{"id": user.ID}
	_, err := u.collection.UpdateOne(context.TODO(), filter, bson.M{"$set": user})
	return err
}

func (u *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"email": email}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)
	return &user, err
}

func (u *UserRepository) GetUserByPhoneNumber(phoneNumber string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"phone_number": phoneNumber}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)
	return &user, err
}

func (u *UserRepository) GetUserByAnonymousDifferentiator(differentiator string) (*domain.User, error) {
	var anonUser domain.User
	filter := bson.M{"differentiator": differentiator}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&anonUser)
	return &anonUser, err
}

func (u *UserRepository) GetUserByID(id string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"id": id}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)
	return &user, err
}
