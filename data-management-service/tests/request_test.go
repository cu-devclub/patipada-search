package tests

import (
	"bytes"
	"data-management/communication"
	"data-management/config"
	"data-management/database"
	"data-management/logging"
	"data-management/request/handlers"
	"data-management/request/repositories"
	"data-management/request/usecases"
	validator "data-management/structValidator"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func setUpTestEnvironment() handlers.Handlers {
	logging.NewSLogger()
	config.InitializeViper("../")
	config.ReadConfig()
	cfg := config.GetConfig()
	db, _ := database.NewMongoDatabase(&cfg)
	v := validator.NewValidator()
	grpc, _ := communication.NewgRPC(&cfg)
	rabbit, err := communication.ConnectToRabbitMQ(&cfg)
	if err != nil {
		log.Println("Error connecting to RabbitMQ", err)
		return nil
	}

	comm := communication.NewCommunicationImpl(*grpc, *rabbit)

	requestRepositories := repositories.NewRequestRepositories(db.GetDb(), &comm)

	requestUsecase := usecases.NewRequestUsecase(&requestRepositories, &v)

	requestHandlers := handlers.NewRequestHandler(&requestUsecase)

	return requestHandlers
}

func createNewRequestFromRawJSON(httpMethod string, url string, payload string) *http.Request {
	req, _ := http.NewRequest(httpMethod, url, bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
	token := generateToken()
	req.Header.Set("Authorization", token)
	return req
}

func createNewRequestFromDataType(httpMethod string, url string, payload any) *http.Request {
	p, _ := json.Marshal(payload)
	req, _ := http.NewRequest(httpMethod, url, bytes.NewBuffer(p))
	req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
	token := generateToken()
	req.Header.Set("Authorization", token)
	return req
}

func generateToken() string {
	// ! Just for development purpose; credentials in production will be different
	reqBody := map[string]string{
		"username": "super-admin",
		"password": "super-admin",
	}
	payload, _ := json.Marshal(reqBody)
	// making request to localhost:8082/login
	req, _ := http.NewRequest("POST", "http://localhost:8082/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json") // Set the Content-Type header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read http response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert response body to string
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	token := response["token"]
	return token.(string)
}
