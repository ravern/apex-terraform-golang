package main

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	apex "github.com/apex/go-apex"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
	"github.com/pkg/errors"
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
type Counter struct {
	UserIP    string `json:"user_ip"`
	Timestamp string `json:"timestamp"`
	Value     int    `json:"value"`
}

// Response JSON
type response struct {
	StatusCode int       `json:"statusCode"`
	Body       []Counter `json:"body"`
}

func main() {
	log.SetOutput(os.Stderr)
	apex.HandleFunc(handle)
}

func handle(evt json.RawMessage, ctx *apex.Context) (interface{}, error) {
	// Unmarshal the JSON
	var e event
	if err := json.Unmarshal(evt, &e); err != nil {
		return newErrorResponse(errors.New("Integer 'value' is required as query")), nil
	}

	// Extract parameters
	ctr := extract(e)

	// Put it in DB
	if err := put(ctr); err != nil {
		return newErrorResponse(err), nil
	}

	// Construct response
	return newSuccessResponse(), nil
}

func extract(e event) Counter {
	ip := strings.Split(e.Headers.XForwardedFor, ",")[0]
	t := time.Now().UTC().Format(time.RFC3339)
	val := e.QueryStringParameters.Value
	return Counter{ip, t, val}
}

func put(c Counter) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return errors.Wrap(err, "Could not load AWS config")
	}

	db := dynamodb.New(cfg)
	table := os.Getenv("DYNAMODB_COUNTER")

	item, err := dynamodbattribute.MarshalMap(c)
	if err != nil {
		panic(err)
	}

	in := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(table),
	}
	req := db.PutItemRequest(in)
	_, err = req.Send()
	if err != nil {
		return errors.Wrap(err, "DynamoDB ain't cooperating")
	}

	return nil
}

func newErrorResponse(err error) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": 500,
		"body":       `{"error":"` + err.Error() + `"}`,
	}
}

func newSuccessResponse() map[string]interface{} {
	return map[string]interface{}{
		"statusCode": 200,
		"body":       `{"hello":"world"}`,
	}
}
