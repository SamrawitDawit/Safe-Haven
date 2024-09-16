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

func (u *UserRepository) CreateUser(user *domain.User) *domain.CustomError {
	_, err := u.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return domain.ErrUserCreationFailed
	}
	return nil
}

func (u *UserRepository) UpdateUser(user *domain.User) *domain.CustomError {
	user.UpdatedAt = time.Now()
	fmt.Println(user.ID)
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	_, err := u.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.ErrUserUpdateFailed
	}
	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (*domain.User, *domain.CustomError) {
	var user domain.User
	filter := bson.M{"email": email}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, domain.ErrUserNotFound
	} else {
		return &user, nil
	}
}

func (u *UserRepository) GetUserByPhoneNumber(phoneNumber string) (*domain.User, *domain.CustomError) {
	var user domain.User
	filter := bson.M{"phoneNumber": phoneNumber}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	return &user, nil
}

func (u *UserRepository) GetUserByAnonymousDifferentiator(differentiator string) (*domain.User, *domain.CustomError) {
	var anonUser domain.User
	filter := bson.M{"anonymousDifferentiator": differentiator}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&anonUser)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}
	return &anonUser, nil
}

func (u *UserRepository) GetUserByID(id uuid.UUID) (*domain.User, *domain.CustomError) {
	var user domain.User
	filter := bson.M{"_id": id}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}
	user.ID = id
	return &user, nil
}

func (u *UserRepository) GetUsersCount() (int, *domain.CustomError) {
	count, err := u.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return 0, domain.ErrUserCountFailed
	}
	return int(count), nil
}
