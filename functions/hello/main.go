package main

import (
	"encoding/json"

	apex "github.com/apex/go-apex"
)

func main() {
	apex.HandleFunc(handle)
}

func handle(evt json.RawMessage, ctx *apex.Context) (interface{}, error) {
	return map[string]interface{}{
		"statusCode": 200,
		"body":       `{ "hello": "world" }`,
	}, nil
}
