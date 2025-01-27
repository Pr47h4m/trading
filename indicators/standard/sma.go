package standard

import (
	"fmt"
	"github.com/pr47h4m/trading"
)

/*
SMASeries - Simple Moving Average
Calculates the entire SMA series for the provided candles and period.
*/
func SMASeries(candles []trading.Candle, period int) ([]trading.SMAValue, error) {
	if len(candles) < period {
		return nil, fmt.Errorf("not enough data or invalid period")
	}

	sma := make([]trading.SMAValue, len(candles)-period+1)
	var sum float64

	for i, c := range candles {
		sum += c.C // Add the close price
		if i >= period-1 {
			if i >= period {
				sum -= candles[i-period].C // Subtract the oldest candle in the period
			}
			sma[i-period+1].Time = c.T
			sma[i-period+1].Value = sum / float64(period) // Calculate SMA
		}
	}

	return sma, nil
}
