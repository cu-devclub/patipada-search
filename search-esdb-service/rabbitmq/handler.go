package rabbitmq

import (
	"encoding/json"
	"log/slog"
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
		var model models.UpdateRecord

		// Convert the map to JSON
		data, err := json.Marshal(payload.Data)
		if err != nil {
			slog.Error("Receive event queue update record; Failed to marshal payload data:", slog.Any("error", err))
			return
		}

		// Unmarshal the JSON into the model
		err = json.Unmarshal(data, &model)
		if err != nil {
			slog.Error("Receive event queue update record; Failed to unmarshal payload data:", slog.Any("error", err))
			return
		}
		consumer.updateRecordEvent(model)

	default:
		slog.Warn("Unknown payload name(event type)", slog.Any("type", payload.Name))
		return
	}
}

func (consumer *Consumer) updateRecordEvent(model models.UpdateRecord) {
	slog.Info("Receive update record event queue....",
		slog.Any("model", model.ToString()),
	)
	err := consumer.recordUsecase.UpdateRecord(&model)
	if err != nil {
		return
	}
}
