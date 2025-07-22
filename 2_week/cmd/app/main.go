package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"project/config"
	"project/internal/db"
	"project/internal/models"
	"project/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("ошибка загрузки конфигурации: %v", err)
	}

	dbConn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("ошибка подключения к БД: %v", err)
	}
	defer dbConn.Close()

	if err := db.MigrateUp(dbConn); err != nil {
		log.Fatalf("ошибка миграции: %v", err)
	}

	ctx := context.Background()
	txService := service.NewTransactionService(dbConn)

	// Демонстрация: создадим транзакцию
	newTx := models.Transaction{
		ID:            12345,
		TerminalUUID:  "terminal-001",
		OrderID:       "order-12345",
		Amount:        1000.50,
		Status:        "PENDING",
		CreatedAt:     time.Now(),
		StatusChanged: time.Now(),
		Code:          "INIT",
		Message:       "Transaction initialized",
	}

	if err := txService.CreateTransaction(ctx, newTx); err != nil {
		log.Fatalf("ошибка создания транзакции: %v", err)
	}
	fmt.Println("Транзакция успешно создана")

	// Получим все транзакции
	allTxs, err := txService.GetAllTransactions(ctx)
	if err != nil {
		log.Fatalf("ошибка получения всех транзакций: %v", err)
	}

	fmt.Println("\nСписок всех транзакций:")
	for _, tx := range allTxs {
		fmt.Printf("- ID: %s, Статус: %s, Сумма: %.2f\n", tx.ID, tx.Status, tx.Amount)
	}

	// Получим транзакции по статусу
	status := "PENDING"
	statusTxs, err := txService.GetTransactionsByStatus(ctx, status)
	if err != nil {
		log.Fatalf("ошибка получения транзакций по статусу: %v", err)
	}

	fmt.Printf("\nТранзакции со статусом %s:\n", status)
	for _, tx := range statusTxs {
		fmt.Printf("- ID: %s, Сумма: %.2f, Сообщение: %s\n", tx.ID, tx.Amount, tx.Message)
	}
}
