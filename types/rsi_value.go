package types

import (
	"encoding/json"
	"time"
)

type RSIValue struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

func (v RSIValue) String() string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}
