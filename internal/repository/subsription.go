package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

// SubscriptionRepository wraps a SQL database.
type SubscriptionRepository struct {
	DB *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{DB: db}
}

// Create inserts a new subscription record with a generated token.
func (r *SubscriptionRepository) Create(email, city, token string) error {
	_, err := r.DB.Exec(
		`INSERT INTO subscriptions (email, city, token, confirmed, unsubscribed, created_at)
         VALUES (?, ?, ?, 0, 0, ?)`,
		email, city, token, time.Now(),
	)
	return err
}

// Confirm sets the confirmed flag for the subscription with a matching token.
func (r *SubscriptionRepository) Confirm(token string) (bool, error) {
	res, err := r.DB.Exec(
		"UPDATE subscriptions SET confirmed = 1 WHERE token = ?", token,
	)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	return count > 0, err
}

// Unsubscribe marks a subscription as unsubscribed by token.
func (r *SubscriptionRepository) Unsubscribe(token string) (bool, error) {
	res, err := r.DB.Exec(
		"UPDATE subscriptions SET unsubscribed = 1 WHERE token = ?", token,
	)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	return count > 0, err
}
