package models

import "time"

type LinkRecord struct {
	ID        int64     `json:"id"`
	Short     string    `json:"short"`
	Source    string    `json:"source"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewLinkRecord struct {
	Short  string `json:"short"`
	Source string `json:"source"`
}
