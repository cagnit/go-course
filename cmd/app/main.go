package main

import (
	"fmt"
	"log"

	"github.com/justifyzz/go-course/internal/calc"
	"github.com/shopspring/decimal"
)

func main() {
	fmt.Print("Введите сумму: ")
	var input float64
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal("Ошибка ввода:", err)
	}

	amount := decimal.NewFromFloat(input)
	result := calc.Multiply(amount)
	fmt.Println(result, "KZT")
}
