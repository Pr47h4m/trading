package trading

import "time"

type Candler interface {
	GetTime() time.Time
	GetOpen() float64
	GetHigh() float64
	GetLow() float64
	GetClose() float64
}

type SMAValue struct {
	Time  time.Time
	Value float64
}

type EMAValue struct {
	Time  time.Time
	Value float64
}

type RSIValue struct {
	Time  time.Time
	Value float64
}

type DMIValue struct {
	Time                 time.Time
	PlusDI, MinusDI, ADX float64
}

type MACDValue struct {
	Time time.Time
}
