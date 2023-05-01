package responses

import "time"

type GetTodoResponse struct {
	ID               int       `json:"id,omitempty"`
	ActitvityGroupID int       `json:"activity_group_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	IsActive         bool      `json:"is_active,omitempty"`
	Priority         string    `json:"priority,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}
