package database

import (
	"database/sql"
	"time"
)

// Password represents a password entry
type Password struct {
	ID        int       `json:"id"`
	Account   string    `json:"account"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// PasswordService handles password operations
type PasswordService struct {
	DB *sql.DB
}

// NewPasswordService creates a new password service
func NewPasswordService(db *sql.DB) *PasswordService {
	return &PasswordService{DB: db}
}

// CreatePassword creates a new password entry
func (ps *PasswordService) CreatePassword(account, username, password string) error {
	query := `INSERT INTO passwords (account, username, password) VALUES (?, ?, ?)`
	_, err := ps.DB.Exec(query, account, username, password)
	return err
}

// GetAllPasswords retrieves all password entries
func (ps *PasswordService) GetAllPasswords() ([]*Password, error) {
	query := `SELECT id, account, username, password, created_at FROM passwords ORDER BY created_at DESC`
	rows, err := ps.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passwords []*Password
	for rows.Next() {
		p := &Password{}
		err := rows.Scan(&p.ID, &p.Account, &p.Username, &p.Password, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, p)
	}

	return passwords, nil
}

// GetPasswordByID retrieves a password entry by ID
func (ps *PasswordService) GetPasswordByID(id int) (*Password, error) {
	query := `SELECT id, account, username, password, created_at FROM passwords WHERE id = ?`
	row := ps.DB.QueryRow(query, id)

	p := &Password{}
	err := row.Scan(&p.ID, &p.Account, &p.Username, &p.Password, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// UpdatePassword updates an existing password entry
func (ps *PasswordService) UpdatePassword(id int, account, username, password string) error {
	query := `UPDATE passwords SET account = ?, username = ?, password = ? WHERE id = ?`
	_, err := ps.DB.Exec(query, account, username, password, id)
	return err
}

// DeletePassword deletes a password entry
func (ps *PasswordService) DeletePassword(id int) error {
	query := `DELETE FROM passwords WHERE id = ?`
	_, err := ps.DB.Exec(query, id)
	return err
}
