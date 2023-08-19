package schema

type CreateRequest struct {
	Author string `json:"author" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

type CreateResponse struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
