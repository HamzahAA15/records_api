package records

import (
	"context"
	"fmt"

	"github.com/HamzahAA15/records_api/utils/errorMap"
)

type RecordsService interface {
	GetFilteredRecords(ctx context.Context, req RequestPayload) ([]Record, error)
}

type recordsService struct {
	recordRepository RecordsRepository
}

func NewRecordsService(recordRepository RecordsRepository) RecordsService {
	return &recordsService{
		recordRepository: recordRepository,
	}
}

func (r *recordsService) GetFilteredRecords(ctx context.Context, req RequestPayload) ([]Record, error) {
	records := []Record{}

	recordsModel, err := r.recordRepository.GetAllRecordsByDate(context.Background(), req.StartDate, req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get all records: %w", err)
	}

	if len(recordsModel) == 0 {
		return nil, fmt.Errorf("%w within the date range", errorMap.RecordsNotFound)
	}

	for _, recordModel := range recordsModel {
		//calculate total marks for each person
		totalMarks := sumMarks(recordModel.Marks)
		if totalMarks < req.MinCount || totalMarks > req.MaxCount {
			continue
		}

		record := Record{
			ID:         recordModel.ID,
			CreatedAt:  recordModel.CreatedAt,
			TotalMarks: totalMarks,
		}

		records = append(records, record)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("%w with given count criteria", errorMap.RecordsNotFound)
	}

	return records, nil
}

func sumMarks(marks []int64) int64 {
	var sum int64
	for _, mark := range marks {
		sum += mark
	}

	return sum
}
