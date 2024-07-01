package service

import (
	"github.com/PetkovaDiana/shop/internal/repository"
	"github.com/PetkovaDiana/shop/internal/service/models"
)

type category struct {
	repo repository.Category
}

func NewCategory(repo repository.Category) Category {
	return &category{repo: repo}
}

func (c *category) GetAllCategory() ([]models.GetAllCategories, error) {
	return c.repo.GetAllCategories()
}
