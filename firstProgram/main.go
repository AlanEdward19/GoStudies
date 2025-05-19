package main

import "fmt"

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	finalAmount := calculateInvestmentReturn(investmentAmount, expectedReturnRate, years)
	futureRealValue := calculateInflationAdjustedReturn(finalAmount, inflationRate, years)

	fmt.Printf("The final amount after %f years is: $%.2f\n", years, finalAmount)
	fmt.Printf("The inflation-adjusted value of the investment is: $%.2f\n", futureRealValue)
}
