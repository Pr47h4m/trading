package types

import (
	"encoding/json"
	"time"
)

type EMAValue struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

func (v EMAValue) String() string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}
