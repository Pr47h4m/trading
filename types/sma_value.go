package types

import (
	"encoding/json"
	"time"
)

type SMAValue struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

func (v SMAValue) String() string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}
