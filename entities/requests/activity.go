package requests

type CreateActivityRequest struct {
	Title string `json:"title" form:"title" query:"title"`
	Email string `json:"email" form:"email" query:"email" validate:"email"`
}

type UpdateActivityRequest struct {
	Title string `json:"title" form:"title" query:"title"`
}
