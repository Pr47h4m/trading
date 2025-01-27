package standard

import (
	"fmt"
	"github.com/pr47h4m/trading"
)

func EMASeries(candles []trading.Candle, period int) ([]trading.EMAValue, error) {
	if len(candles) <= period {
		return nil, fmt.Errorf("not enough data or invalid period")
	}

	ema := make([]trading.EMAValue, len(candles)-period+1)
	multiplier := 2.0 / float64(period+1)

	var smaSum float64
	for i := 0; i < period; i++ {
		smaSum += candles[i].C
	}
	initialEMA := smaSum / float64(period)
	ema[0] = trading.EMAValue{Time: candles[period-1].T, Value: initialEMA}

	for i := period; i < len(candles); i++ {
		currentEMA := (candles[i].C * multiplier) + (ema[i-period].Value * (1 - multiplier))
		ema[i-period+1] = trading.EMAValue{Time: candles[i].T, Value: currentEMA}
	}

	return ema, nil
}
