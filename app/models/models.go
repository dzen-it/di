package models

import "encoding/json"

type Node struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Address struct {
	Address string `json:"address,omitempty"`
}

type Result struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func NewResult(result, err string) []byte {
	r := Result{
		Result: result,
		Error:  err,
	}

	b, _ := json.Marshal(r)
	return b
}
