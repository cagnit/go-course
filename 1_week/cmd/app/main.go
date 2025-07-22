package main

import (
	"fmt"
	"project/internal/input"
	"project/internal/weekly"
	"project/pkg/convert"
)

func main() {

	amount := input.GetDecimalInput("USD>")
	result := convert.Convert_to_USD(amount)
	fmt.Println(result.String() + " KZT")
	weekly.WeeklyReport()
}
