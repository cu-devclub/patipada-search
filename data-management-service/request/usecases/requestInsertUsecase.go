package usecases

import (
	"data-management/errors"
	"data-management/messages"
	"data-management/request/entities"
	"data-management/request/helper"
	"data-management/request/models"
	"log"
	"time"
)

// InsertRequest inserts a new Request into the database.
// It takes a pointer to a models.Request as an argument.
// The function will return an error if the insertion fails, otherwise it will return nil.
//
// The function first validates the Request index using the ValidateRequestIndex method of the RequestRepositories.
// If the validation fails, it returns an error.
//
// Then, it generates a unique requestID for the new Request using the helper.GeneraterequestID function.
// If the generation of the requestID fails, it returns an internal server error.
//
// The function then populates a entities.Request struct with the data from the models.Request struct,
// the generated requestID, and the current time for the CreatedAt and UpdatedAt fields.
// The Status field is set to "pending".
//
// Then, it inserts the entities.Request struct into the database using the InsertRequest method of the RequestRepositories.
// If the insertion fails, it returns an error.
//
// Finally, it populates the original models.Request struct with the requestID, CreatedAt, UpdatedAt, and Status fields from the entities.Request struct.
//
// Error Status Codes:
//
//	400: ERR_REQUEST_INDEX_NOT_EXISTS
//	500: INTERNAL_SERVER_ERROR
//
// Usage:
//
//	request := &models.Request{...}
//	err := requestUsecase.InsertRequest(Request)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *requestUsecase) InsertRequest(request *models.Request) *errors.RequestError {
	log.Println("Insert request; Request: ", request)
	if err := r.validator.Validate(request); err != nil {
		log.Println("Error validate request; Request: ", request, "Error: ", err)
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// validate record index
	result,err := r.requestRepositories.ValidateRecordIndex(request.Index); 
	if err != nil || result == false {
		log.Println("Error validate record index; Request: ", request, "Error: ", err)
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	result,err = r.requestRepositories.ValidateUsername(request.By); 
	if err != nil || result == false {
		log.Println("Error validate username; Request: ", request, "Error: ", err)
		return errors.CreateError(400, messages.BAD_REQUEST)
	}


	// populate request entity
	requestID, err := helper.GenerateRequestID(r.requestRepositories)
	if err != nil {
		log.Println("Error generate request ID; Error: ", err)
		return errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	requestEntity := entities.Request{
		RequestID:  requestID,
		Index:      request.Index,
		YoutubeURL: request.YoutubeURL,
		Question:   request.Question,
		Answer:     request.Answer,
		StartTime:  request.StartTime,
		EndTime:    request.EndTime,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Status:     "pending",
		By:         request.By,
		ApprovedBy: "",
	}

	// insert request entity
	objId, err := r.requestRepositories.InsertRequest(&requestEntity)
	if err != nil {
		log.Println("Error insert request entity; RequestEntity: ", requestEntity, "Error: ", err)
		return errors.CreateError(500, messages.ERR_INSERT_REQUEST)
	}

	// populate request model with request entity to return
	request.ID = objId
	request.RequestID = requestEntity.RequestID
	request.CreatedAt = requestEntity.CreatedAt
	request.UpdatedAt = requestEntity.UpdatedAt
	request.Status = requestEntity.Status

	log.Println("Success insert request")
	return nil
}
