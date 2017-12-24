package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	apex "github.com/apex/go-apex"
	log "github.com/sirupsen/logrus"
)

// Event JSON
type event struct {
	QueryStringParameters struct {
		Value int `json:"value"`
	} `json:"queryStringParameters"`
	Headers struct {
		XForwardedFor string `json:"X-Forwarded-For"`
	}
}

// Counter model
type counter struct {
	UserIP    string
	Timestamp time.Time
	Value     int
}

func main() {
	log.SetOutput(os.Stderr)
	apex.HandleFunc(handle)
}

func handle(evt json.RawMessage, ctx *apex.Context) (interface{}, error) {
	// Unmarshal the JSON
	var e event
	if err := json.Unmarshal(evt, &e); err != nil {
		return nil, errors.New("Integer 'value' is required.")
	}

	// Extract parameters
	_ = extract(e)

	// Construct response
	return response(evt), nil
}

func extract(e event) counter {
	return counter{}
}

func response(evt json.RawMessage) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": 200,
		"body":       string(evt),
	}
}
