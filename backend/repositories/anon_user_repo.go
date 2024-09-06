package repositories

import (
	"backend/domain"
	"backend/usecases/interfaces"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnonymousUserRepository struct {
	collection *mongo.Collection
}

func NewAnonUserRepo(db *mongo.Database, collectionName string) interfaces.AnonymousUserRepoInterface {
	return &AnonymousUserRepository{
		collection: db.Collection(collectionName),
	}
}

func (a *AnonymousUserRepository) CreateAnonymousUser(anonUser *domain.AnonymousUser) error {
	_, err := a.collection.InsertOne(context.TODO(), anonUser)
	return err
}

func (a *AnonymousUserRepository) GetUserByAnonymousDifferentiator(differentiator string) (*domain.AnonymousUser, error) {
	var anonUser domain.AnonymousUser
	filter := bson.M{"differentiator": differentiator}
	err := a.collection.FindOne(context.TODO(), filter).Decode(&anonUser)
	return &anonUser, err
}
