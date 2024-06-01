package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 6.5

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Time Period (in years): ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
	fmt.Println(futureValue - futureRealValue)
}
