package converter

import "github.com/shopspring/decimal"

const rate = 520

func CalculateSalary(userNumber int64) decimal.Decimal {
	userDecimal := decimal.NewFromInt(userNumber)
    rateDecimal := decimal.NewFromInt(rate)
    return userDecimal.Mul(rateDecimal)
}
