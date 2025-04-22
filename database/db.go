package database

import (
	"callback-info/models"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

const MaxRecords = 10000

func InitDB(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	// Create table if not exists
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS callbacks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ip TEXT,
			domain TEXT,
			params TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}

func SaveCallback(ip, domain, params string) error {
	// Start a transaction
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert new record
	_, err = tx.Exec("INSERT INTO callbacks (ip, domain, params) VALUES (?, ?, ?)",
		ip, domain, params)
	if err != nil {
		return err
	}

	// Check total count
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM callbacks").Scan(&count)
	if err != nil {
		return err
	}

	// If count exceeds MaxRecords, delete oldest records
	if count > MaxRecords {
		_, err = tx.Exec(`
			DELETE FROM callbacks 
			WHERE id IN (
				SELECT id FROM callbacks 
				ORDER BY created_at ASC 
				LIMIT ?
			)
		`, count-MaxRecords)
		if err != nil {
			return err
		}
	}

	// Commit transaction
	return tx.Commit()
}

func GetCallbacks(page, pageSize int) (models.CallbackResponse, error) {
	var response models.CallbackResponse
	var total int64

	// Get total count (will be at most MaxRecords)
	err := DB.QueryRow("SELECT COUNT(*) FROM callbacks").Scan(&total)
	if err != nil {
		return response, err
	}

	// Get callbacks
	query := "SELECT id, ip, domain, params, created_at FROM callbacks ORDER BY created_at DESC LIMIT ? OFFSET ?"
	offset := (page - 1) * pageSize
	rows, err := DB.Query(query, pageSize, offset)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	var callbacks []models.Callback
	for rows.Next() {
		var c models.Callback
		err := rows.Scan(&c.ID, &c.IP, &c.Domain, &c.Params, &c.CreatedAt)
		if err != nil {
			return response, err
		}
		callbacks = append(callbacks, c)
	}

	response.Total = total
	response.Callbacks = callbacks
	return response, nil
}
