package repositories

import (
	"backend/domain"
	"backend/usecases/interfaces"
	"backend/utils"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CaseRepository struct {
	collection *mongo.Collection
}

func NewCaseRepo(db *mongo.Database, collectionName string) interfaces.CaseRepositoryInterface {
	return &CaseRepository{
		collection: db.Collection(collectionName),
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
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": CaseID}, bson.M{"$set": fields})
	if err != nil {
		utils.LogError("Failed to update Case", err)
		return domain.ErrCaseUpdateFailed
	}
	return nil
}

func (r *CaseRepository) GetCaseByID(CaseID uuid.UUID) (*domain.Case, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var Case *domain.Case
	err := r.collection.FindOne(ctx, bson.M{"_id": CaseID}).Decode(&Case)
	if err != nil {
		utils.LogError("Failed to get Case", err)
		return nil, domain.ErrCaseNotFound
	}
	return Case, nil
}

func (r *CaseRepository) GetCasesBySubmitterID(SubmitterID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var Cases []*domain.Case
	cursor, err := r.collection.Find(ctx, bson.M{"SubmitterID": SubmitterID})
	if err != nil {
		utils.LogError("Failed to get Cases", err)
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

func (r *CaseRepository) GetCasesByCounselorID(counselorID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var Cases []*domain.Case
	cursor, err := r.collection.Find(ctx, bson.M{"counselor_id": counselorID})
	if err != nil {
		utils.LogError("Failed to get Cases", err)
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
	if err != nil {
		utils.LogError("Failed to get Cases", err)
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
	var Cases []*domain.Case
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		utils.LogError("Failed to get Cases", err)
		return nil, domain.ErrCaseFetchFailed
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
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
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": CaseID})
	if err != nil {
		utils.LogError("Failed to delete Case", err)
		return domain.ErrCaseUpdateFailed
	}
	return nil
}
