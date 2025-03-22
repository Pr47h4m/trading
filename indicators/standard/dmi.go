package standard

import (
	"fmt"
	"math"

	"github.com/pr47h4m/trading/types"
)

func DMISeries(candles []*types.Candle, period int) ([]types.DMIValue, error) {
	if len(candles) <= period {
		return nil, fmt.Errorf("not enough data to calculate DMI")
	}

	n := len(candles)

	tr := make([]float64, n)
	plusDM := make([]float64, n)
	minusDM := make([]float64, n)
	dmi := make([]types.DMIValue, len(candles))

	// Calculate True Range (TR), +DM, and -DM
	for i := 1; i < len(candles); i++ {
		up := candles[i].H - candles[i-1].H
		down := candles[i-1].L - candles[i].L

		// True Range
		tr[i] = math.Max(candles[i].H-candles[i].L,
			math.Max(math.Abs(candles[i].H-candles[i-1].C), math.Abs(candles[i].L-candles[i-1].C)))

		// Positive Directional Movement (+DM)
		if up > down && up > 0 {
			plusDM[i] = up
		}

		// Negative Directional Movement (-DM)
		if down > up && down > 0 {
			minusDM[i] = down
		}
	}

	// rma calculates the Rolling Moving Average (Wilder's Smoothing) for the given data and period.
	rma := func(data []float64, period int) []float64 {
		rma := make([]float64, len(data))
		sum := 0.0

		for i := 0; i < len(data); i++ {
			if i < period {
				sum += data[i]
				if i == period-1 {
					rma[i] = sum / float64(period)
				}
			} else {
				rma[i] = (rma[i-1]*(float64(period)-1) + data[i]) / float64(period)
			}
		}

		return rma
	}

	// Smooth TR, +DM, and -DM using Wilder's smoothing (RMA - Rolling Moving Average)
	smoothedTR := rma(tr, period)
	smoothedPlusDM := rma(plusDM, period)
	smoothedMinusDM := rma(minusDM, period)

	// Calculate +DI and -DI
	for i := period; i < len(candles); i++ {
		positiveDI := (smoothedPlusDM[i] / smoothedTR[i]) * 100
		negativeDI := (smoothedMinusDM[i] / smoothedTR[i]) * 100
		dmi[i].PlusDI = positiveDI
		dmi[i].MinusDI = negativeDI
	}

	// Calculate ADX
	dx := make([]float64, len(candles))
	for i := period; i < len(candles); i++ {
		plus := dmi[i].PlusDI
		minus := dmi[i].MinusDI
		dx[i] = math.Abs(plus-minus) / math.Max(plus+minus, 1) * 100 // Avoid division by zero
	}

	smoothedDX := rma(dx, period)
	for i := period; i < len(candles); i++ {
		dmi[i].ADX = smoothedDX[i]
	}

	return dmi, nil
}
