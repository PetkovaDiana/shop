package repository

import (
	"context"
	domainModels "github.com/PetkovaDiana/shop/internal/service/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Category interface {
	GetAllCategories(ctx context.Context, categoriesIDs []int64) ([]*domainModels.GetAllCategories, error)
}

type Product interface {
	GetAllProduct(ctx context.Context, productIDs []int64) ([]*domainModels.GetAllProducts, error)
}

type Authorization interface {
	CreateClient(ctx context.Context, client domainModels.CreateClient) error
	GetClient(ctx context.Context, email string) (*domainModels.Client, error)
}

type Repo struct {
	Category
	Product
	Authorization
}

func NewRepository(pool *pgxpool.Pool) *Repo {
	return &Repo{
		Category:      NewCategoryRepo(pool),
		Product:       NewProductRepo(pool),
		Authorization: NewAuthRepo(pool),
	}
}
