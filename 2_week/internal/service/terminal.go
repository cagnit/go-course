package service

import (
	"context"
	"database/sql"
	"fmt"

	"project/internal/models"
)

type TerminalService struct {
	DB *sql.DB
}

func NewTerminalService(db *sql.DB) *TerminalService {
	return &TerminalService{DB: db}
}

// CreateTerminal сохраняет новый терминал в БД
func (s *TerminalService) CreateTerminal(ctx context.Context, t models.Terminal) error {
	query := `INSERT INTO terminals (id, client_id, client_secret, uuid) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.ExecContext(ctx, query, t.ID, t.ClientID, t.ClientSecret, t.UUID)
	if err != nil {
		return fmt.Errorf("CreateTerminal: %w", err)
	}
	return nil
}

// GetAllTerminals возвращает список всех терминалов
func (s *TerminalService) GetAllTerminals(ctx context.Context) ([]models.Terminal, error) {
	query := `SELECT id, client_id, client_secret, uuid FROM terminals`

	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("GetAllTerminals: %w", err)
	}
	defer rows.Close()

	var terminals []models.Terminal
	for rows.Next() {
		var t models.Terminal
		if err := rows.Scan(&t.ID, &t.ClientID, &t.ClientSecret, &t.UUID); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		terminals = append(terminals, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return terminals, nil
}
