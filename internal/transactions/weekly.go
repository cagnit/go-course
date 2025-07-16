package transactions

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func ProcessWeekly(transactions []map[int]decimal.Decimal) {
	days := map[int]string{
		1: "Понедельник",
		2: "Вторник",
		3: "Среда",
		4: "Четверг",
		5: "Пятница",
		6: "Суббота",
		7: "Воскресенье",
	}

	// Суммы по дням недели
	dayTotals := make(map[int]decimal.Decimal)
	totalWeek := decimal.NewFromInt(0)

	for _, transaction := range transactions {
		for day, amount := range transaction {
			dayTotals[day] = dayTotals[day].Add(amount)
			totalWeek = totalWeek.Add(amount)
		}
	}

	for i := 1; i <= 7; i++ {
		fmt.Println(days[i])

		sum := dayTotals[i]

		if sum.GreaterThan(decimal.Zero) {
			fmt.Println("Поступление:", sum)
		} else if sum.LessThan(decimal.Zero) {
			fmt.Println("Списание:", sum.Abs())
		} else {
			fmt.Println("Операций не было")
		}

		fmt.Println()
	}

	fmt.Println("Итог за неделю:", totalWeek)
}
