package usecases

import (
	"log/slog"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
)

func (r *recordUsecaseImpl) UpdateRecord(record *models.UpdateRecord) error {
	//TODO : We have to generate new vector for each record if its update
	updateRecordEntity := helper.UpdateRecordModelToEntity(record)
	if err := r.recordRepository.UpdateRecord(updateRecordEntity); err != nil {
		slog.Error("Failed to update record",
			slog.String("Record", record.ToString()),
			slog.String("err", err.Error()),
		)
		return err
	}

	slog.Info("Update record successfully", slog.String("Record", record.DocumentID))

	return nil
}
