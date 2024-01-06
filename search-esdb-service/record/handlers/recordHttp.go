package handlers

import (
	"search-esdb-service/record/usecases"
)

type recordHttpHandler struct {
	recordUsecase usecases.RecordUsecase
}

func NewRecordHttpHandler(recordUsecase usecases.RecordUsecase) RecordHandler {
	return &recordHttpHandler{
		recordUsecase: recordUsecase,
	}
}
