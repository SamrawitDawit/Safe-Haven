package repositories

import (
	"backend/domain"
	"backend/usecases/interfaces"
	"context"

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


func (u *UserRepository) CreateNormalUser(user *domain.User) error {
	_, err := u.collection.InsertOne(context.TODO(), user)
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

