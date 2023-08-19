package model

import "time"

type Article struct {
	Id      int64     `json:"id" gorm:"column:id"`
	Author  string    `json:"author" gorm:"column:author"`
	Title   string    `json:"title" gorm:"column:title"`
	Body    string    `json:"body" gorm:"column:body"`
	Created time.Time `json:"created" gorm:"column:created"`
}

func (a *Article) TableName() string {
	return "article"
}
