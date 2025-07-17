package main

import (
    "fmt"
	"go-course/internal/converter"
    "go-course/internal/transaction"
)

func main() {
    for {
        fmt.Println("1 — Конвертация")
        fmt.Println("2 — Расчёт транзакций за неделю")
        fmt.Println("0 — Выход")
        fmt.Print("Выберите действие: ")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            var userNumber int64
    		fmt.Print("Введите количество часов: ")
    		fmt.Scan(&userNumber)

    		salary := converter.CalculateSalary(userNumber)
    		fmt.Printf("Результат: %s KZT\n\n", salary.String())
        case 2:
            transaction.ProcessTransactions()
        case 0:
            fmt.Println("Выход.")
            return
        default:
            fmt.Println("Неизвестная команда.\n")
        }
    }
}
