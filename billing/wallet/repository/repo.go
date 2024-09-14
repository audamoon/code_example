package repository

import (
	"billing_service/internal/http_errors"
	"database/sql"
	"fmt"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetCurrentBalance(tx *sql.Tx, walletID int) (int, error) {
	query := `
		select balance
		from billing.wallets
		where id = $1;
    `

	var balance int

	if err := tx.QueryRow(query, walletID).Scan(&balance); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("Repository.GetCurrentBalance: %w", http_errors.ErrNotFound)
		}

		return 0, fmt.Errorf("Repository.GetCurrentBalance: %w", err)
	}

	return balance, nil
}

func (r *Repository) UpdateBalance(tx *sql.Tx, walletID, balance int) error {
	query := `
		update billing.wallets 
		set balance = $1
		where id = $2
    `

	if _, err := tx.Exec(query, balance, walletID); err != nil {
		return fmt.Errorf("Repository.UpdateBalance #1: %w", err)
	}

	return nil
}
