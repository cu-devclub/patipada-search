package event

import (
	"encoding/json"
	"log"
	"search-esdb-service/constant"
	"search-esdb-service/record/models"
)

type Payload struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func (consumer *Consumer) handlePayload(payload Payload) {
	switch payload.Name {
	case constant.UPDATE_RECORD_PAYLOAD_NAME:
		log.Println("payload Name :", payload.Name, "payload Data :", payload.Data)
		var model models.UpdateRecord

		// Convert the map to JSON
		data, err := json.Marshal(payload.Data)
		if err != nil {
			log.Println("Failed to marshal payload data:", err)
			return
		}

		// Unmarshal the JSON into the model
		err = json.Unmarshal(data, &model)
		if err != nil {
			log.Println("Failed to unmarshal payload data:", err)
			return
		}
		consumer.updateRecordEvent(model)

	default:
		log.Println("Unknown event type")
		return
	}
}

func (consumer *Consumer) updateRecordEvent(model models.UpdateRecord) {
	err := consumer.recordUsecase.UpdateRecord(&model)
	if err != nil {
		log.Println("Failed to update record:", err)
		return
	}

	log.Println("Record updated successfully")
}
