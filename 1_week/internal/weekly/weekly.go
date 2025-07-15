package weekly

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func WeeklyReport() {
	transactions := []map[int]decimal.Decimal{
		{1: decimal.NewFromFloat(25000.00)},
		{1: decimal.NewFromFloat(-25000.00)},
		{2: decimal.NewFromFloat(25000.00)},
		{3: decimal.NewFromFloat(-25000.00)},
		{4: decimal.NewFromFloat(25000.00)},
		{5: decimal.NewFromFloat(-25000.00)},
		{6: decimal.NewFromFloat(-25000.00)},
		{7: decimal.NewFromFloat(-25000.00)},
	}

	dayNames := map[int]string{
		1: "Понедельник",
		2: "Вторник",
		3: "Среда",
		4: "Четверг",
		5: "Пятница",
		6: "Суббота",
		7: "Воскресенье",
	}

	allDays := make(map[int]decimal.Decimal)
	Summary := decimal.Zero

	// Суммируем все транзакции по дням
	for _, m := range transactions {
		for day, amount := range m {
			allDays[day] = allDays[day].Add(amount)
			Summary = Summary.Add(amount)
		}
	}

	// Выводим результаты
	for day := 1; day <= 7; day++ {
		amount, ok := allDays[day]
		if !ok {
			continue
		}

		dayName := dayNames[day]
		fmt.Println(dayName)

		// здесь if / else if / else
		if amount.GreaterThan(decimal.Zero) {
			fmt.Printf("Поступление: %s\n\n", amount.StringFixed(2))
		} else if amount.LessThan(decimal.Zero) {
			fmt.Printf("Списание: %s\n\n", amount.Abs().StringFixed(2))
		} else {
			fmt.Println("Нет операций")
		}

	}
	switch {
	case Summary.GreaterThan(decimal.Zero):
		fmt.Printf("Итог за неделю: %s (мы в плюсе)", Summary.String())
	case Summary.LessThan(decimal.Zero):
		fmt.Printf("Итог за неделю: %s (мы в минусе)", Summary.String())
	default:
		fmt.Print("Выйти в ноль тоже победа)")
	}
}
