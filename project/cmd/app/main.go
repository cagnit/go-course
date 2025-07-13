package main

import (
	"fmt"
	"project/internal/convert"
	"project/internal/input"
)

func main() {
	amount := input.GetDecimalInput("USD>")
	result := convert.Convert_to_USD(amount)
	fmt.Printf(result.String() + " KZT")
}
