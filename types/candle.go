package types

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
