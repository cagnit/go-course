package service

import (
	"context"
	"database/sql"
	"fmt"

	"project/internal/models"
)

type TransactionService struct {
	DB *sql.DB
}

func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{DB: db}
}

// CreateTransaction сохраняет новую транзакцию
func (s *TransactionService) CreateTransaction(ctx context.Context, tx models.Transaction) error {
	query := `INSERT INTO transactions (
		id, terminal_id, order_id, amount, status,
		created_at, status_changed, code, message
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	_, err := s.DB.ExecContext(ctx, query,
		tx.ID, tx.TerminalUUID, tx.OrderID, tx.Amount, tx.Status,
		tx.CreatedAt, tx.StatusChanged, tx.Code, tx.Message,
	)

	if err != nil {
		return fmt.Errorf("CreateTransaction: %w", err)
	}
	return nil
}

// GetAllTransactions возвращает все транзакции
func (s *TransactionService) GetAllTransactions(ctx context.Context) ([]models.Transaction, error) {
	query := `SELECT id, terminal_id, order_id, amount, status,
		created_at, status_changed, code, message FROM transactions`

	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("GetAllTransactions: %w", err)
	}
	defer rows.Close()

	var txs []models.Transaction
	for rows.Next() {
		var tx models.Transaction
		err := rows.Scan(
			&tx.ID, &tx.TerminalUUID, &tx.OrderID, &tx.Amount, &tx.Status,
			&tx.CreatedAt, &tx.StatusChanged, &tx.Code, &tx.Message,
		)
		if err != nil {
			return nil, fmt.Errorf("Scan: %w", err)
		}
		txs = append(txs, tx)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return txs, nil
}

// GetTransactionsByStatus возвращает транзакции по статусу
func (s *TransactionService) GetTransactionsByStatus(ctx context.Context, status string) ([]models.Transaction, error) {
	query := `SELECT id, terminal_id, order_id, amount, status,
		created_at, status_changed, code, message FROM transactions WHERE status = $1`

	rows, err := s.DB.QueryContext(ctx, query, status)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionsByStatus: %w", err)
	}
	defer rows.Close()

	var txs []models.Transaction
	for rows.Next() {
		var tx models.Transaction
		err := rows.Scan(
			&tx.ID, &tx.TerminalUUID, &tx.OrderID, &tx.Amount, &tx.Status,
			&tx.CreatedAt, &tx.StatusChanged, &tx.Code, &tx.Message,
		)
		if err != nil {
			return nil, fmt.Errorf("Scan: %w", err)
		}
		txs = append(txs, tx)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return txs, nil
}
