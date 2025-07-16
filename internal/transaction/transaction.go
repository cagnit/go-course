package transaction

import (
    "fmt"
    "github.com/shopspring/decimal"
)

var dayNames = map[int]string{
    1: "Понедельник",
    2: "Вторник",
    3: "Среда",
    4: "Четверг",
    5: "Пятница",
    6: "Суббота",
    7: "Воскресенье",
}

func ProcessTransactions() {
    transactions := []map[int]decimal.Decimal{
        {1: decimal.NewFromFloat(25000.00)},
        {1: decimal.NewFromFloat(20000.00)},
        {2: decimal.NewFromFloat(-9800.00)},
        {3: decimal.NewFromFloat(-1222.22)},
        {4: decimal.NewFromFloat(-1500.07)},
        {5: decimal.NewFromFloat(1201.37)},
        {6: decimal.NewFromFloat(-100.32)},
        {7: decimal.NewFromFloat(-523.33)},
    }

    weeklyTotalsByDay := make(map[int]decimal.Decimal)
    var weeklyGrandTotal decimal.Decimal

    for _, transactionMap := range transactions {
        for dayOfWeek, transactionAmount := range transactionMap {
            if currentTotalForDay, dayAlreadyExists := weeklyTotalsByDay[dayOfWeek]; dayAlreadyExists {
                weeklyTotalsByDay[dayOfWeek] = currentTotalForDay.Add(transactionAmount)
            } else {
                weeklyTotalsByDay[dayOfWeek] = transactionAmount
            }
        }
    }

    for dayOfWeek := 1; dayOfWeek <= 7; dayOfWeek++ {
        totalForDay := weeklyTotalsByDay[dayOfWeek]
        weeklyGrandTotal = weeklyGrandTotal.Add(totalForDay)

        fmt.Println(dayNames[dayOfWeek])
        switch {
        case totalForDay.GreaterThan(decimal.Zero):
            fmt.Printf("Поступление: %s\n\n", totalForDay.StringFixed(2))
        case totalForDay.LessThan(decimal.Zero):
            fmt.Printf("Списание: %s\n\n", totalForDay.Abs().StringFixed(2))
        default:
            fmt.Println("Нет транзакций\n")
        }
    }

    fmt.Printf("Итог за неделю: %s\n", weeklyGrandTotal.StringFixed(2))
}
