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

type ReportRepository struct {
	collection *mongo.Collection
}

func NewReportRepo(db *mongo.Database, collectionName string) interfaces.ReportRepositoryInterface {
	return &ReportRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *ReportRepository) CreateReport(report *domain.Report) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, report)
	if err != nil {
		utils.LogError("Failed to create report", err)
		return domain.ErrReportCreationFailed
	}
	return nil
}

func (r *ReportRepository) UpdateReportFields(reportID uuid.UUID, fields map[string]interface{}) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": reportID}, bson.M{"$set": fields})
	if err != nil {
		utils.LogError("Failed to update report", err)
		return domain.ErrReportUpdateFailed
	}
	return nil
}

func (r *ReportRepository) GetReportByID(reportID uuid.UUID) (*domain.Report, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var report *domain.Report
	err := r.collection.FindOne(ctx, bson.M{"_id": reportID}).Decode(&report)
	if err != nil {
		utils.LogError("Failed to get report", err)
		return nil, domain.ErrReportNotFound
	}
	return report, nil
}

func (r *ReportRepository) GetReportsByReporterID(reporterID uuid.UUID) ([]*domain.Report, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var reports []*domain.Report
	cursor, err := r.collection.Find(ctx, bson.M{"reporter_id": reporterID})
	if err != nil {
		utils.LogError("Failed to get reports", err)
		return nil, domain.ErrReportFetchFailed
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var report *domain.Report
		if err := cursor.Decode(&report); err != nil {
			utils.LogError("Failed to decode report", err)
			return nil, domain.ErrReportFetchFailed
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func (r *ReportRepository) GetReportsByCounselorID(counselorID uuid.UUID) ([]*domain.Report, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var reports []*domain.Report
	cursor, err := r.collection.Find(ctx, bson.M{"counselor_id": counselorID})
	if err != nil {
		utils.LogError("Failed to get reports", err)
		return nil, domain.ErrReportFetchFailed
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var report *domain.Report
		if err := cursor.Decode(&report); err != nil {
			utils.LogError("Failed to decode report", err)
			return nil, domain.ErrReportFetchFailed
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func (r *ReportRepository) GetReportsByStatus(status string) ([]*domain.Report, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	var reports []*domain.Report
	cursor, err := r.collection.Find(ctx, bson.M{"status": status})
	if err != nil {
		utils.LogError("Failed to get reports", err)
		return nil, domain.ErrReportFetchFailed
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var report *domain.Report
		if err := cursor.Decode(&report); err != nil {
			utils.LogError("Failed to decode report", err)
			return nil, domain.ErrReportFetchFailed
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func (r *ReportRepository) GetAllReports() ([]*domain.Report, *domain.CustomError) {
	var reports []*domain.Report
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		utils.LogError("Failed to get reports", err)
		return nil, domain.ErrReportFetchFailed
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var report *domain.Report
		if err := cursor.Decode(&report); err != nil {
			utils.LogError("Failed to decode report", err)
			return nil, domain.ErrReportFetchFailed
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func (r *ReportRepository) DeleteReport(reportID uuid.UUID) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": reportID})
	if err != nil {
		utils.LogError("Failed to delete report", err)
		return domain.ErrReportUpdateFailed
	}
	return nil
}
