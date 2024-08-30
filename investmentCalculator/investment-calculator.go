package investmentCalculator

import (
	"fmt"
	"math"
	"strconv"
)

func investmentCalculator() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter Expected Rate of Return: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Enter Number of Years Invested: ")
	fmt.Scan(&years)

	futureValue := calculateFV(investmentAmount, expectedReturnRate, years)
	futureRealValue := calculatRealFV(futureValue, years)

	outputText(floatToString(futureValue))
	outputText(floatToString(futureRealValue))
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFV(investmentAmount, expectedReturnRate, years float64) float64 {
	return investmentAmount * math.Pow(1+expectedReturnRate/100, years)
}

func calculatRealFV(futureValue, years float64) float64 {
	const inflationRate = 2.5

	return futureValue / math.Pow(1+inflationRate/100, years)
}

func floatToString(num float64) string {
	return strconv.FormatFloat(num, 'E', -1, 64)
}
