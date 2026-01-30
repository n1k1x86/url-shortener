package models

import "time"

type LinkRecord struct {
	ID        int64
	Short     string
	Source    string
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewLinkRecord struct {
	Short  string `json:"short"`
	Source string `json:"source"`
}
