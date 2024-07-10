package service

import (
	"context"
	"github.com/PetkovaDiana/shop/internal/repository"
	"github.com/PetkovaDiana/shop/internal/service/models"
)

type product struct {
	repo repository.Product
	ctx  context.Context
}

func NewProduct(ctx context.Context, repo repository.Product) Product {
	return &product{ctx: ctx, repo: repo}
}

func (p *product) GetProduct(ctx context.Context, filter models.GetProductsFilter) ([]*models.GetAllProducts, error) {
	return p.repo.GetAllProduct(ctx, filter.ProductsIDs)
}
