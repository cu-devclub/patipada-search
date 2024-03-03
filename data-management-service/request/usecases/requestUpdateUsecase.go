package usecases

import (
	"data-management/errors"
	"data-management/messages"
	"data-management/request/helper"
	"data-management/request/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateRequest updates a request in the MongoDB collection.
//
// Return
// - Success : no error (nil)
// - 400 : Bad request (validation error)
// - 500 : internal server error
//
// The function takes a pointer to a models.Request object as input. The Request object is first validated
// and then converted to an entity using the helper.ModelsToEntity function. The UpdatedAt field of the entity
// is set to the current time.
//
// The function then calls the UpdateRequest method of the requestRepositories to update the document in the
// MongoDB collection. If the update operation fails, the function returns an error. If the update operation
// is successful, the function returns nil.
//
// The function also checks if the record index of the request is valid by calling the ValidateRecordIndex
// method of the requestRepositories. If the record index is not valid, the function returns an error.
//
// Example:
//
//	request := &models.Request{
//	    ID: "60d5ecf7c88f9a200f9e2c5a",
//	    Question: "Updated question",
//	    ... other fields ...
//	}
//	err := r.UpdateRequest(request)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *requestUsecase) UpdateRequest(request *models.Request) *errors.RequestError {
	log.Println("Update request usecase; Request: ", request)
	// validate the request

	log.Println("Update request usecase; Validating request ....")
	if err := r.validator.Validate(request); err != nil {
		log.Println("Error validate request; Request: ", request, "Error: ", err)
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// check valid record index
	log.Println("Update request usecase; Validating record index ....")
	result, err := r.requestRepositories.ValidateRecordIndex(request.Index)
	if err != nil || result == false {
		log.Println("Error validate record index; Request: ", request, "Error: ", err)
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// check valid by
	log.Println("Update request usecase; Validating username ....")
	result, err = r.requestRepositories.ValidateUsername(request.By)
	if err != nil || result == false {
		log.Println("Error validate username; Request: ", request, "Error: ", err)
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// check valid approved by
	log.Println("Update request usecase; Validating approved by ....")
	result, err = r.requestRepositories.ValidateUsername(request.ApprovedBy)
	if err != nil || result == false {
		log.Println("Error validate approved by; Request: ", request, "Error: ", err)
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// ---  Get all requests that has the same record Index by Get usecase
	log.Println("Update request usecase; Getting all requests ....")
	requests, er := r.GetRequest("", "", "", request.Index, "")
	if er != nil {
		log.Println("Error get request; Request: ", request, "Error: ", er, "ERR == nil", er != nil)
		return errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	// ---  Update all request that come before the current request by setting status to "reviewed"
	for _, req := range requests {
		if req.UpdatedAt.After(request.UpdatedAt) {
			continue
		}

		if req.ID == request.ID {
			continue
		}

		log.Println("Updating usecase ; update old request; Request: ", req)
		req.Status = "reviewed"
		requestEntitiy := helper.ModelsToEntity(req)
		requestEntitiy.UpdatedAt = time.Now()
		if err := r.requestRepositories.UpdateRequest(requestEntitiy); err != nil {
			log.Println("Error update request; Request: ", request, "Error: ", err)
			return errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
		}
	}

	// --- Update the current request
	log.Println("Update request usecase; Updating request ....")
	requestEntity := helper.ModelsToEntity(request)
	requestEntity.UpdatedAt = time.Now()
	requestEntity.Status = "reviewed"
	if err := r.requestRepositories.UpdateRequest(requestEntity); err != nil {
		log.Println("Error update request; Request: ", request, "Error: ", err)
		if err == mongo.ErrNoDocuments {
			return errors.CreateError(400, messages.BAD_REQUEST)
		}
		return errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	recordEntity := helper.RequestToRecordsEntity(requestEntity)
	recordEntity.ExtractHTML()

	// --- Update the record by sending plain text of `startTime`,`endTime`,`question`,`answer` to the record (search) service
	log.Println("Update request usecase; Updating record ....")
	err = r.requestRepositories.UpdateRecord(recordEntity)
	if err != nil {
		log.Println("Error update record; Request: ", request, "Error: ", err)
		return errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	return nil
}
