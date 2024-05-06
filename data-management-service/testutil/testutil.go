package testutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"

)

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	return r.MatchString(uuid)
}

func CreateNewRequest(httpMethod string, url string, payload any) *http.Request {
	p, _ := json.Marshal(payload)
	req, _ := http.NewRequest(httpMethod, url, bytes.NewBuffer(p))
	req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
	token := "mock-token"
	req.Header.Set("Authorization", token)
	return req
}

