package service

import (
	"context"
	"github.com/PetkovaDiana/shop/internal/repository"
	"github.com/PetkovaDiana/shop/internal/service/models"
)

type category struct {
	repo repository.Category
	ctx  context.Context
}

func NewCategory(ctx context.Context, repo repository.Category) Category {
	return &category{
		ctx:  ctx,
		repo: repo,
	}
}

func (c *category) GetCategory(ctx context.Context, filter models.GetCategoriesFilter) ([]*models.GetAllCategories, error) {
	return c.repo.GetAllCategories(ctx, filter.CategoriesIDs)
}
