package models

import "github.com/google/uuid"

// Flag represents a feature flag resource within an environment
type Flag struct {
	ID uuid.UUID `json:"id" db:"id"`
	EnvironmentID uuid.UUID `json:"environment_id" db:"environment_id"`
	Key         string    `json:"key" db:"key"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Enabled     bool      `json:"enabled" db:"enabled"`
	Variations  []Variations  `json:"variations" db:"variations"`
	DefaultValue string    `json:"default_value" db:"default_value"`
	CreatedAt   string    `json:"created_at" db:"created_at"`
	UpdatedAt   string    `json:"updated_at" db:"updated_at"`
}

// Variations represents the possible variations for a feature flag
type Variations struct {
	ID 		uuid.UUID `json:"id" db:"id"`
	Value 	string    `json:"value" db:"value"`
	Name 	string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
}