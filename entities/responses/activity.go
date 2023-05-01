package responses

import "time"

type GetActivityResponse struct {
	ID        int       `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
