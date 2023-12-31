package schema

import "time"

type GetResponse struct {
	ID      int64     `json:"id" `
	Author  string    `json:"author" `
	Title   string    `json:"title" `
	Body    string    `json:"body" `
	Created time.Time `json:"created" `
}

type MessageConsume struct {
	ID      int64     `json:"id" `
	Author  string    `json:"author" `
	Title   string    `json:"title" `
	Body    string    `json:"body" `
	Created time.Time `json:"created" `
}
