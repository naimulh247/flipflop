package models

import "github.com/google/uuid"

// Project represents a project resource
type Project struct {
	ID           uuid.UUID     `json:"id" db:"id"`
	Key          string        `json:"key" db:"key"`
	Name         string        `json:"name" db:"name"`
	Description  string        `json:"description" db:"description"`
	CreatedAt    string        `json:"created_at" db:"created_at"`
	UpdatedAt    string        `json:"updated_at" db:"updated_at"`
	Environments []Environment `json:"environments,omitempty" db:"environments"`
}
