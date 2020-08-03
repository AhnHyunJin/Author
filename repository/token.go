package repo

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/promptech1/infuser-author/model"
)

type TokenRepository interface {
	FindByToken(token string) *model.Token
}

type tokenRepositoryDB struct {
	DB *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepositoryDB{DB: db}
}

func (r *tokenRepositoryDB) FindByToken(token string) *model.Token {
	t := model.Token{}

	if r.DB.Where("token = ?", token).First(&t).RecordNotFound() {
		return nil
	}

	return &t
}