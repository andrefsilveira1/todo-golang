package models

import "time"

type Data struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Completed   bool      `json:"completed,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"CreatedAt,omitempty"`
}
