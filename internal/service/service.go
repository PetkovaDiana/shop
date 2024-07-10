package service

import (
	"context"
	"github.com/PetkovaDiana/shop/internal/repository"
	"github.com/PetkovaDiana/shop/internal/service/models"
)

type Category interface {
	GetCategory(ctx context.Context, filter models.GetCategoriesFilter) ([]*models.GetAllCategories, error)
}

type Product interface {
	GetProduct(ctx context.Context, filter models.GetProductsFilter) ([]*models.GetAllProducts, error)
}

type Authorization interface {
	CreateClient(ctx context.Context, client models.CreateClient) error
	AuthClient(ctx context.Context, client models.AuthClient) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Category
	Product
	Authorization
}

func NewService(ctx context.Context, repo *repository.Repo, passwordHash ItemsArgon) *Service {
	return &Service{
		Category:      NewCategory(ctx, repo.Category),
		Product:       NewProduct(ctx, repo.Product),
		Authorization: NewAuthService(ctx, repo.Authorization, passwordHash),
	}
}
