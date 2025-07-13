package input

import (
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

func GetDecimalInput(prompt string) decimal.Decimal {
	var a string
	fmt.Print(prompt)
	fmt.Scan(&a)
	b, err := decimal.NewFromString(a)
	if err != nil {
		log.Fatalf("Неверный формат: %v", err)
	}
	return b
}
