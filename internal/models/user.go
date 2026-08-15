package models

import "time"

type Organization struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"` // nama perusahaan/workspace, mis. "Gradivus" atau nama klien SaaS
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organization_id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"-"`
	Role           string    `json:"role"` // owner | admin | member
	CreatedAt      time.Time `json:"created_at"`
}
