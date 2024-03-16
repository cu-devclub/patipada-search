package usecases

import (
	"log/slog"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
)

func (r *recordUsecaseImpl) UpdateRecord(record *models.UpdateRecord) error {
	updateRecordEntity := helper.UpdateRecordModelToEntity(record)
	if err := r.recordRepository.UpdateRecord(updateRecordEntity); err != nil {
		slog.Error("Failed to update record",
			slog.String("Record", record.ToString()),
			slog.String("err", err.Error()),
		)
		return err
	}

	slog.Info("Update record successfully", slog.String("Record", record.ToString()))

	return nil
}
