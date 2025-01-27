package standard

import (
	"fmt"
	"github.com/pr47h4m/trading"
)

/*
SMASeries - Simple Moving Average
*/
func SMASeries(candles []trading.Candler, period int) ([]trading.SMAValue, error) {
	if len(candles) <= period {
		return nil, fmt.Errorf("not enough data or invalid period")
	}

	sma := make([]trading.SMAValue, len(candles)-period+1)

	var sum float64

	for i, c := range candles {
		sum += c.GetClose()
		if i >= period-1 {
			if i >= period {
				sum -= candles[i-period].GetClose()
			}
			sma[i-period+1].Time = c.GetTime()
			sma[i-period+1].Value = sum / float64(period)
		}
	}

	return nil, nil
}

/*
LastSMA - Simple Moving Average

Note: Work in Progress
*/
func LastSMA(candles []trading.Candler, period int) (*trading.SMAValue, error) {

	return nil, nil
}

/*
NextSMA - Simple Moving Average

Note: Work in Progress
*/
func NextSMA(series []trading.SMAValue, candle trading.Candler, period int) (*trading.SMAValue, error) {

	return nil, nil
}
