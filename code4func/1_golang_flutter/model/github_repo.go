package model

import "time"

type GithubRepo struct {
	Name string `json:"name" db:"name,omitempty"`
	Description string `json:"description" db:"description,omitempty"`
	Url string `json:"url" db:"url,omitempty"`
	Color string `json:"color" db:"color,omitempty"`
	Lang string `json:"lang" db:"lang,omitempty"`
	Fork string `json:"fork" db:"fork,omitempty"`
	Starts string `json:"stars" db:"stars,omitempty"`
	StartsToday string `json:"stars_today" db:"stars_today,omitempty"`
	BuildBy string `json:"build_by" db:"build_by,omitempty"`
	CreatedAt time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at,omitempty"`
}