package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println(futureValue)
}
