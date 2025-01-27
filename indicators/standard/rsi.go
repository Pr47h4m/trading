package standard

import (
	"fmt"
	"github.com/pr47h4m/trading"
)

func RSISeries(candles []trading.Candler, period int) ([]trading.RSIValue, error) {
	if len(candles) <= period {
		return nil, fmt.Errorf("not enough data or invalid period")
	}

	rsi := make([]trading.RSIValue, len(candles)-period)

	gains, losses := 0.0, 0.0

	for i := 1; i < period; i++ {
		change := candles[i].GetClose() - candles[i-1].GetClose()
		if change > 0 {
			gains += change
		} else {
			losses -= change
		}
	}

	avgGain := gains / float64(period)
	avgLoss := losses / float64(period)

	rsi[0].Time = candles[period].GetTime()
	rsi[0].Value = 100 - (100 / (1 + avgGain/avgLoss))

	for i := period + 1; i < len(candles); i++ {
		change := candles[i].GetClose() - candles[i-1].GetClose()
		if change > 0 {
			gains = change
			losses = 0
		} else {
			gains = 0
			losses = -change
		}

		avgGain = ((avgGain * float64(period-1)) + gains) / float64(period)
		avgLoss = ((avgLoss * float64(period-1)) + losses) / float64(period)

		rsi[i-period].Time = candles[i].GetTime()
		if avgLoss == 0 {
			rsi[i-period].Value = 100
		} else {
			rs := avgGain / avgLoss
			rsi[i-period].Value = 100 - (100 / (1 + rs))
		}

	}

	return rsi, nil
}

func LastRSI(candles []trading.Candler, period int) (*trading.RSIValue, error) {

	return nil, nil
}

func NextRSI(series []trading.RSIValue, candle trading.Candler, period int) (*trading.RSIValue, error) {

	return nil, nil
}
