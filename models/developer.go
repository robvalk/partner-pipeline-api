package models

import (
	"time"

	"github.com/google/uuid"
)

// Developer defines model for Developer.
type Developer struct {

	// The timestamp of when this resource was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	ID        *uuid.UUID `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`

	// List of skills that the Developer has
	Skils *[]string `json:"skils,omitempty"`
	Team  *string   `json:"team,omitempty"`

	// The timestamp of when this resource was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewDeveloper creates a new instance of Developer
func NewDeveloper(name, team *string, skills *[]string) *Developer {
	newID, _ := uuid.NewRandom()

	return &Developer{
		CreatedAt: &time.Time{},
		ID:        &newID,
		Name:      name,
		Skils:     skills,
		Team:      team,
		UpdatedAt: &time.Time{},
	}
}
