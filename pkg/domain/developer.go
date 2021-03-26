package domain

import "time"

type Developer struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Team      string    `json:"team,omitempty"`
	Skills    []string  `json:"skils,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
