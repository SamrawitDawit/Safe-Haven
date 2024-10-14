package repositories

import (
	"backend/domain"
	"backend/usecases/interfaces"
	"backend/utils"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CaseRepository struct {
	collection *mongo.Collection
}

func NewCaseRepo(db *mongo.Database, collectionName string) interfaces.CaseRepositoryInterface {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	collection := db.Collection(collectionName)
	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"submitter_id": 1}, // 1 for ascending order
		Options: options.Index(),
	})
	if err != nil {
		// Handle error
		utils.LogError("Error creating index for submitter_id:", err)
		return nil
	}

	// Repeat for counselor_id and status
	_, err = collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"counselor_id": 1},
		Options: options.Index(),
	})
	if err != nil {
		utils.LogError("Error creating index for counselor_id:", err)
		return nil
	}

	_, err = collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"status": 1},
		Options: options.Index(),
	})
	if err != nil {
		utils.LogError("Error creating index for status:", err)
		return nil
	}

	return &CaseRepository{
		collection: collection,
	}
}

func (cr *CaseRepository) CreateCase(Case *domain.Case) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	_, err := cr.collection.InsertOne(ctx, Case)
	if err != nil {
		utils.LogError("Failed to create Case", err)
		return domain.ErrCaseCreationFailed
	}
	return nil
}

func (r *CaseRepository) UpdateCaseFields(CaseID uuid.UUID, fields map[string]interface{}) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": CaseID}, bson.M{"$set": fields})
	if err != nil {
		utils.LogError("Failed to update Case", err)
		return domain.ErrCaseUpdateFailed
	}
	if result.ModifiedCount == 0 {
		return domain.ErrCaseNotFound
	}

	return nil
}

func (r *CaseRepository) GetCaseByID(CaseID uuid.UUID) (*domain.Case, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var Case *domain.Case
	err := r.collection.FindOne(ctx, bson.M{"_id": CaseID}).Decode(&Case)
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrCaseNotFound
	} else if err != nil {
		utils.LogError("Failed to get Case", err)
		return nil, domain.ErrCaseFetchFailed
	}
	return Case, nil
}

func (r *CaseRepository) GetCasesBySubmitterID(SubmitterID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var Cases []*domain.Case
	cursor, err := r.collection.Find(ctx, bson.M{"submitter_id": SubmitterID})
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrCaseNotFound
	} else if err != nil {
		utils.LogError("Failed to get Case", err)
		return nil, domain.ErrCaseFetchFailed
	}
	defer cursor.Close(ctx)
	if cursor.Err() != nil {
		utils.LogError("Failed to get Cases", cursor.Err())
		return nil, domain.ErrCaseFetchFailed
	}

	for cursor.Next(ctx) {
		var Case *domain.Case
		if err := cursor.Decode(&Case); err != nil {
			utils.LogError("Failed to decode Case", err)
			return nil, domain.ErrCaseFetchFailed
		}
		Cases = append(Cases, Case)
	}
	return Cases, nil
}

func (r *CaseRepository) GetCasesByCounselorID(counselorID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var Cases []*domain.Case
	cursor, err := r.collection.Find(ctx, bson.M{"counselor_id": counselorID})
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrCaseNotFound
	} else if err != nil {
		utils.LogError("Failed to get Case", err)
		return nil, domain.ErrCaseFetchFailed
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var Case *domain.Case
		if err := cursor.Decode(&Case); err != nil {
			utils.LogError("Failed to decode Case", err)
			return nil, domain.ErrCaseFetchFailed
		}
		Cases = append(Cases, Case)
	}
	return Cases, nil
}

func (r *CaseRepository) GetCasesByStatus(status string) ([]*domain.Case, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var Cases []*domain.Case
	cursor, err := r.collection.Find(ctx, bson.M{"status": status})
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrCaseNotFound
	} else if err != nil {
		utils.LogError("Failed to get Case", err)
		return nil, domain.ErrCaseFetchFailed
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var Case *domain.Case
		if err := cursor.Decode(&Case); err != nil {
			utils.LogError("Failed to decode Case", err)
			return nil, domain.ErrCaseFetchFailed
		}
		Cases = append(Cases, Case)
	}
	return Cases, nil
}

func (r *CaseRepository) GetAllCases() ([]*domain.Case, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var Cases []*domain.Case
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrCaseNotFound
	} else if err != nil {
		utils.LogError("Failed to get Case", err)
		return nil, domain.ErrCaseFetchFailed
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var Case *domain.Case
		if err := cursor.Decode(&Case); err != nil {
			utils.LogError("Failed to decode Case", err)
			return nil, domain.ErrCaseFetchFailed
		}
		Cases = append(Cases, Case)
	}
	return Cases, nil
}

func (r *CaseRepository) DeleteCase(CaseID uuid.UUID) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": CaseID})
	if err != nil {
		utils.LogError("Failed to delete Case", err)
		return domain.ErrCaseUpdateFailed
	}
	if result.DeletedCount == 0 {
		return domain.ErrCaseNotFound
	}
	return nil
}
