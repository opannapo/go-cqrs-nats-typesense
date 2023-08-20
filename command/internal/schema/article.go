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

type MessagePublish struct {
	ID      int64  `json:"id" `
	Author  string `json:"author" `
	Title   string `json:"title" `
	Body    string `json:"body" `
	Created string `json:"created" `
}
