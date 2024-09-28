package repository

import (
	"database/sql"
	"errors"
)

type RDSRepo struct {
	db *sql.DB
}

func NewRDSRepo(db *sql.DB) *RDSRepo {
	return &RDSRepo{
		db: db,
	}
}

func (r *RDSRepo) Put(key, value string) error {
	_, err := r.db.Exec(`INSERT INTO kv (key, value) VALUES (?, ?)
                         ON CONFLICT(key) DO UPDATE SET value = excluded.value;`,
		key, value)
	if err != nil {
		return err
	}

	return err
}

func (r *RDSRepo) Get(key string) (string, bool, error) {
	var value string
	err := r.db.QueryRow(`SELECT value FROM kv WHERE key = ?;`, key).Scan(&value)
	if errors.Is(err, sql.ErrNoRows) {
		return "", false, nil // Key does not exist
	} else if err != nil {
		return "", false, err
	}

	return value, true, nil
}
