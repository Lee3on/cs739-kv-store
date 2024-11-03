package repository

import (
	"cs739-kv-store/models"
	"database/sql"
	"encoding/json"
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
	_, err := r.db.Exec(`INSERT INTO kv (Key, Value) VALUES (?, ?)
                         ON CONFLICT(Key) DO UPDATE SET Value = excluded.Value;`,
		key, value)
	if err != nil {
		return err
	}

	return err
}

func (r *RDSRepo) Get(key string) (string, bool, error) {
	var value string
	err := r.db.QueryRow(`SELECT Value FROM kv WHERE Key = ?;`, key).Scan(&value)
	if errors.Is(err, sql.ErrNoRows) {
		return "", false, nil // Key does not exist
	} else if err != nil {
		return "", false, err
	}

	return value, true, nil
}

// Serialize all data to JSON
func (r *RDSRepo) Serialize() ([]byte, error) {
	rows, err := r.db.Query(`SELECT Key, Value FROM kv`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kvPairs []models.KVPair
	for rows.Next() {
		var kv models.KVPair
		if err := rows.Scan(&kv.Key, &kv.Value); err != nil {
			return nil, err
		}
		kvPairs = append(kvPairs, kv)
	}

	return json.Marshal(kvPairs)
}

// Deserialize JSON data and load it into the database
func (r *RDSRepo) Deserialize(data []byte) error {
	var kvPairs []models.KVPair
	if err := json.Unmarshal(data, &kvPairs); err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO kv (Key, Value) VALUES (?, ?)
                             ON CONFLICT(Key) DO UPDATE SET Value = excluded.Value;`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, kv := range kvPairs {
		if _, err := stmt.Exec(kv.Key, kv.Value); err != nil {
			return err
		}
	}

	return tx.Commit()
}
