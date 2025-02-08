package types

import (
	"encoding/json"
	"time"
)

type DMIValue struct {
	Time    time.Time `json:"time"`
	PlusDI  float64   `json:"plusDI"`
	MinusDI float64   `json:"MinusDI"`
	ADX     float64   `json:"adx"`
}

func (v DMIValue) String() string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}
