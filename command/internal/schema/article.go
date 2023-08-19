package schema

type CreateRequest struct {
	Author string `json:"author" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body"`
}
