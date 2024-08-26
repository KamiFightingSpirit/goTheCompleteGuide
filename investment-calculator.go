package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Println("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter Number of Years Invested: ")
	fmt.Scan(&years)
	fmt.Println("Enter Expected Rate of Return: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
