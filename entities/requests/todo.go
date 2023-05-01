package requests

type CreateTodoRequest struct {
	Title           string `json:"title" form:"title" query:"title"`
	ActivityGroupID int    `json:"activity_group_id" form:"activity_group_id" query:"activity_group_id"`
	IsActive        bool   `json:"is_active" form:"is_active" query:"is_active"`
	Priority        string `json:"priority" form:"priority" query:"priority"`
}

type UpdateTodoRequest struct {
	Title    string `json:"title" form:"title" query:"title"`
	Priority string `json:"priority" form:"priority" query:"priority"`
	IsActive bool   `json:"is_active" form:"is_active" query:"is_active"`
	Status   string `json:"status" form:"status" query:"status"`
}
