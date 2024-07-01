package repository

import (
	domainModels "github.com/PetkovaDiana/shop/internal/service/models"
	"github.com/jmoiron/sqlx"
)

type Category interface {
	GetAllCategories() ([]domainModels.GetAllCategories, error)
}

type Repo struct {
	Category
}

func NewRepository(db *sqlx.DB) *Repo {
	return &Repo{
		Category: NewCategoryRepo(db),
	}
}
