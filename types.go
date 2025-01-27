package trading

import (
	"encoding/json"
	"time"
)

type Candle struct {
	T  time.Time
	O  float64 // The open price for the symbol in the given time period.
	H  float64 // The highest price for the symbol in the given time period.
	L  float64 // The lowest price for the symbol in the given time period.
	C  float64 // The close price for the symbol in the given time period.
	v  int     // The trading volume of the symbol in the given time period.
	vw float64 // The volume weighted average price.
}

func (v Candle) String() string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}

type SMAValue struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

type DMAValue struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

func (v DMAValue) String() string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}

type EMAValue struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

func (v EMAValue) String() string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}

type RSIValue struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

func (v RSIValue) String() string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}

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

type MACDValue struct {
	Time time.Time `json:"time"`
}

func (v MACDValue) String() string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}
