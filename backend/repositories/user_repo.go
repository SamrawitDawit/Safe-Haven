package repositories

import (
	"backend/domain"
	"backend/usecases/interfaces"
	"backend/utils"
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

const queryTimeout = 10 * time.Second

func (u *UserRepository) CreateUser(user *domain.User) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		utils.LogError("Failed to create user", err)
		return domain.ErrUserCreationFailed
	}
	return nil
}

func (u *UserRepository) UpdateUser(user *domain.User) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err := u.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		utils.LogError("Failed to update user", err)
		return domain.ErrUserUpdateFailed
	}
	return nil
}

func (u *UserRepository) UpdateUserFields(userID uuid.UUID, fields map[string]interface{}) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	// Add an automatic "updatedAt" field to the map
	fields["updatedAt"] = time.Now()

	filter := bson.M{"_id": userID}
	update := bson.M{"$set": fields}

	_, err := u.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		utils.LogError("Failed to update fields for user", err)
		return domain.ErrUserUpdateFailed
	}
	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (*domain.User, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	var user domain.User
	filter := bson.M{"email": email}

	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}
		utils.LogError("Error fetching user by email", err)
		return nil, domain.ErrUserFetchFailed
	}
	return &user, nil
}

func (u *UserRepository) GetUserByPhoneNumber(phoneNumber string) (*domain.User, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	var user domain.User
	filter := bson.M{"phoneNumber": phoneNumber}

	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}
		utils.LogError("Error fetching user by phone number", err)
		return nil, domain.ErrUserFetchFailed
	}
	return &user, nil
}

func (u *UserRepository) GetUserByID(id uuid.UUID) (*domain.User, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	var user domain.User
	filter := bson.M{"_id": id}

	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}

		utils.LogError("Error fetching user by ID", err)
		return nil, domain.ErrUserFetchFailed
	}
	return &user, nil
}

func (r *UserRepository) GetUserByIDWithLock(id uuid.UUID) (*domain.User, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var user domain.User

	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println(id)
			return nil, domain.ErrUserNotFound
		}
		utils.LogError("Error fetching user by ID", err)
		return nil, domain.ErrUserFetchFailed
	}
	if user.Lock {
		return nil, domain.ErrRaceCondition
	}
	updated_fields := map[string]interface{}{
		"lock": true,
	}
	cerr := r.UpdateUserFields(id, updated_fields)
	if cerr != nil {
		return nil, cerr
	}

	return &user, nil
}

func (u *UserRepository) GetUsersCount() (int, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	count, err := u.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		utils.LogError("Error counting users", err)
		return 0, domain.ErrUserCountFailed
	}
	return int(count), nil
}
