package service

import (
	"github.com/PetkovaDiana/shop/internal/repository"
	"github.com/PetkovaDiana/shop/internal/service/models"
)

type Category interface {
	GetAllCategory() ([]models.GetAllCategories, error)
}

type Service struct {
	Category
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		Category: NewCategory(repo.Category),
	}
}
