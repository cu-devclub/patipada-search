package tests

import (
	"bytes"
	"data-management/config"
	"data-management/database"
	"data-management/request/handlers"
	"data-management/request/repositories"
	"data-management/request/usecases"
	validator "data-management/structValidator"
	"encoding/json"
	"net/http"
)

func setUpTestEnvironment() handlers.Handlers {
	config.InitializeViper("../")
	cfg := config.GetConfig()
	db := database.NewMongoDatabase(&cfg)
	v := validator.NewValidator()

	requestRepositories := repositories.NewRequestRepositories(db.GetDb())

	requestUsecase := usecases.NewRequestUsecase(requestRepositories, v)

	requestHandlers := handlers.NewRequestHandler(requestUsecase)

	return requestHandlers
}

func createNewRequestFromRawJSON(httpMethod string, url string, payload string) *http.Request {
	req, _ := http.NewRequest(httpMethod, url, bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
	// TODO : implementing token authorization
	return req
}

func createNewRequestFromDataType(httpMethod string, url string, payload any) *http.Request {
	p, _ := json.Marshal(payload)
	req, _ := http.NewRequest(httpMethod, url, bytes.NewBuffer(p))
	req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
	// TODO : implementing token authorization
	return req
}
