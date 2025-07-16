package calc

import (
	"github.com/shopspring/decimal"
)

// Multiply умножает число на фиксированный курс
func Multiply(value decimal.Decimal) decimal.Decimal {
	rate := decimal.NewFromFloat(520.0)
	return value.Mul(rate)
}
