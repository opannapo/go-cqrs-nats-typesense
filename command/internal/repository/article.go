package repository

import (
	"errors"
	"gcnt/internal/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var ArticleRepositoryInstance *ArticleRepositoryImpl

func InitArticleRepositoryInstance() {
	if ArticleRepositoryInstance == nil {
		ArticleRepositoryInstance = &ArticleRepositoryImpl{}
	}
}

type IArticleRepository interface {
	Create(data *model.Article, db *gorm.DB) (err error)
}

type ArticleRepositoryImpl struct {
}

func (a ArticleRepositoryImpl) Create(data *model.Article, db *gorm.DB) (err error) {
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

	return
}
