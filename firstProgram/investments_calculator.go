package main

import "math"

func calculateInvestmentReturn(principal float64, rate float64, years float64) float64 {
	var futureValue = principal * math.Pow(1+rate/100, years)

	return futureValue
}

func calculateInflationAdjustedReturn(futureValue float64, inflationRate float64, years float64) float64 {
	var inflationAdjusted = futureValue / math.Pow(1+inflationRate/100, years)

	return inflationAdjusted
}
