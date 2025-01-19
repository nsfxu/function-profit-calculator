package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = printAndScan("Revenue: ")
	expenses = printAndScan("Expenses: ")
	taxRate = printAndScan("Tax Rate: ")

	ebt, profit, ratio := ebtProfitAndRatioCalculation(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func printAndScan(text string) (value float64) {
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func ebtProfitAndRatioCalculation(revenue float64, expenses float64, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
