package main

import (
	"fmt"
	"math"
)

func interest() {
	const inflationRate = 2
	var investmentAmount float64
	var expectedReturnRate = 5.5
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	var futureRealVal = futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealVal)
}

func profit_calculator() {
	var expenses int
	var revenue int
	var taxRate float64

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter tax Rate: ")
	fmt.Scan(&taxRate)

	var eBt = revenue - expenses
	fmt.Println(eBt)

	var profit = float64(eBt) * (1 - taxRate/100)
	fmt.Println(profit)

	ratio := float64(eBt) / profit
	fmt.Println(ratio)
}

func main() {
	fmt.Println("Choose an option")
	fmt.Println("1. Calculate Profit")
	fmt.Println("2. Calculate Interest")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		profit_calculator()
	case 2:
		interest()
	default:
		fmt.Println("Invalid option")
	}
}
