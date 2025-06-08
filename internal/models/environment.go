package models

import "github.com/google/uuid"

// Environment represents an environment resource within a project
type Environment struct {
	ID          uuid.UUID `json:"id" db:"id"`
	ProjectID   uuid.UUID `json:"project_id" db:"project_id"`
	Key         string    `json:"key" db:"key"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   string    `json:"created_at" db:"created_at"`
	UpdatedAt   string    `json:"updated_at" db:"updated_at"`
}
