package standard

import (
	"fmt"
	"github.com/pr47h4m/trading/types"
)

/*
RSISeries - Relative Strength Index
Calculates the RSI values for the provided candles and period.
*/
func RSISeries(candles []types.Candle, period int) ([]types.RSIValue, error) {
	if len(candles) <= period {
		return nil, fmt.Errorf("not enough data or invalid period")
	}

	rsi := make([]types.RSIValue, len(candles)-period)

	gains, losses := 0.0, 0.0

	for i := 1; i < period; i++ {
		change := candles[i].C - candles[i-1].C
		if change > 0 {
			gains += change
		} else {
			losses -= change
		}
	}

	avgGain := gains / float64(period)
	avgLoss := losses / float64(period)

	rsi[0].Time = candles[period].T
	rsi[0].Value = 100 - (100 / (1 + avgGain/avgLoss))

	for i := period + 1; i < len(candles); i++ {
		change := candles[i].C - candles[i-1].C
		if change > 0 {
			gains = change
			losses = 0
		} else {
			gains = 0
			losses = -change
		}

		avgGain = ((avgGain * float64(period-1)) + gains) / float64(period)
		avgLoss = ((avgLoss * float64(period-1)) + losses) / float64(period)

		rsi[i-period].Time = candles[i].T
		if avgLoss == 0 {
			rsi[i-period].Value = 100
		} else {
			rs := avgGain / avgLoss
			rsi[i-period].Value = 100 - (100 / (1 + rs))
		}
	}

	return rsi, nil
}
