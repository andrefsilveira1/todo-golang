package models

import "time"

type User struct {
	Id        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}
