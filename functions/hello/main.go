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
	UserIP    string    `json:"user_ip" dynamodbav:"UserIP"`
	Timestamp time.Time `json:"timestamp" dynamodbav:"Timestamp"`
	Value     int       `json:"value" dynamodbav:"Value"`
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
		return newErrorResponse(errors.New("Couldn't unmarshal JSON")), nil
	}

	// Extract parameters
	ctr := extract(e)

	// Load AWS stuff
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return newErrorResponse(errors.Wrap(err, "Could not load AWS config")), nil
	}
	db := dynamodb.New(cfg)
	table := os.Getenv("DYNAMODB_COUNTER")

	// Put it in DB
	if err := put(db, table, ctr); err != nil {
		return newErrorResponse(err), nil
	}

	// Read all the counters
	ctrs, err := read(db, table)
	if err != nil {
		return newErrorResponse(err), nil
	}

	// Construct response
	return newSuccessResponse(ctrs), nil
}

func extract(e event) Counter {
	ip := strings.Split(e.Headers.XForwardedFor, ",")[0]
	t := time.Now()
	val := e.QueryStringParameters.Value
	return Counter{ip, t, val}
}

func put(db *dynamodb.DynamoDB, table string, c Counter) error {
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
		return errors.Wrap(err, "Couldn't put item")
	}

	return nil
}

func read(db *dynamodb.DynamoDB, table string) ([]Counter, error) {
	in := &dynamodb.ScanInput{
		TableName: aws.String(table),
	}
	req := db.ScanRequest(in)
	out, err := req.Send()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't scan items")
	}

	ctrs := []Counter{}
	if err := dynamodbattribute.UnmarshalListOfMaps(out.Items, &ctrs); err != nil {
		panic(err)
	}

	return ctrs, nil
}

func newErrorResponse(err error) map[string]interface{} {
	body := map[string]interface{}{
		"error": err.Error(),
	}

	json, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return map[string]interface{}{
		"statusCode": 500,
		"body":       string(json),
	}
}

func newSuccessResponse(ctrs []Counter) map[string]interface{} {
	body := map[string]interface{}{
		"message": "Success!",
		"data":    ctrs,
	}

	json, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return map[string]interface{}{
		"statusCode": 200,
		"body":       string(json),
	}
}
