package convert

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func Convert_to_USD(amount decimal.Decimal) decimal.Decimal {
	curent_rate, _ := decimal.NewFromString("522.37") //можно было через API или из .env
	fmt.Println("Rate from USDC TO KZT:" + curent_rate.String())
	return amount.Mul(curent_rate)
}
