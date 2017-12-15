package main

import (
	"encoding/json"
	"strings"

	apex "github.com/apex/go-apex"
)

type message struct {
	Value string `json:"value"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message

		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		m.Value = strings.ToUpper(m.Value)

		return m, nil
	})
}
