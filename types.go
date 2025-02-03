package trading

import (
	"encoding/json"
	"time"
)

type Candle struct {
	T  time.Time `json:"t" bson:"t"`
	O  float64   `json:"o" bson:"o"`   // The open price for the symbol in the given time period.
	H  float64   `json:"h" bson:"h"`   // The highest price for the symbol in the given time period.
	L  float64   `json:"l" bson:"l"`   // The lowest price for the symbol in the given time period.
	C  float64   `json:"c" bson:"c"`   // The close price for the symbol in the given time period.
	V  int       `json:"v" bson:"v"`   // The trading volume of the symbol in the given time period.
	VW float64   `json:"vw" bson:"vw"` // The volume weighted average price.
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
