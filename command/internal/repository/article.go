package repository

import (
	"context"
	"errors"
	"gcnt/internal/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var ArticleRepositoryInstance IArticleRepository

func InitArticleRepositoryInstance() {
	if ArticleRepositoryInstance == nil {
		ArticleRepositoryInstance = &articleRepository{}
	}
}

type IArticleRepository interface {
	Create(ctx context.Context, data *model.Article, db *gorm.DB) (article model.Article, err error)
}

type articleRepository struct {
}

func (a *articleRepository) Create(ctx context.Context, data *model.Article, db *gorm.DB) (article model.Article, err error) {
	tx := db.Debug().Create(data)
	if tx.Error != nil {
		log.Err(err).Send()
		err = tx.Error
		return
	}

	if tx.RowsAffected < 1 {
		err = errors.New("err failed to create article")
		log.Err(err).Send()
		return
	}

	article = *data
	return
}
