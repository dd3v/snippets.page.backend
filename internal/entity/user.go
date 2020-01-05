package model

import (
	"time"
)

//User represents resources in JSON format.
type User struct {
	ID           int       `json:"_id"`
	Login        string    `json:"login" `
	PasswordHash string    `json:"-" `
	Email        string    `json:"email" `
	Token        string    `json:"token" `
	Banned       bool      `json:"banned" `
	CreatedAt    time.Time `json:"created_at" `
	UpdatedAt    time.Time `json:"updated_at"`
}
