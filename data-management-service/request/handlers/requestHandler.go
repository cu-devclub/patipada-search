package handlers

import "data-management/request/usecases"

type requestHandler struct {
	requestUsecase usecases.UseCase
}

func NewRequestHandler(requestUsecase *usecases.UseCase) Handlers {
	return &requestHandler{
		requestUsecase: *requestUsecase,
	}
}
