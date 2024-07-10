package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/PetkovaDiana/shop/internal/repository/entities"
	domainModels "github.com/PetkovaDiana/shop/internal/service/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type product struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) Product {
	return &product{db: db}
}

func (p *product) GetAllProduct(ctx context.Context, productIDs []int64) ([]*domainModels.GetAllProducts, error) {
	products := make([]*domainModels.GetAllProducts, 0, len(productIDs))

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(
			fmt.Sprintf("p.%s, p.%s, p.%s, p.%s, p.%s",
				entities.Failed_Product_ID,
				entities.Failed_Product_Title,
				entities.Failed_Product_Description,
				entities.Failed_Product_Price,
				entities.Faile_Category_id,
			),
		).
		From(fmt.Sprintf("%s AS p", entities.Table_Product)).
		OrderBy(fmt.Sprintf("p.%s", entities.Failed_Category_ID))

	if len(productIDs) != 0 {
		query = query.Where(squirrel.Eq{fmt.Sprintf("p.%s", entities.Failed_Product_ID): productIDs})
	}

	q, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.db.Query(ctx, q, args...)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("products not found")
	}
	defer rows.Close()

	for rows.Next() {
		productRow := new(domainModels.GetAllProducts)
		err = rows.Scan(&productRow.ID, &productRow.Title, &productRow.Description, &productRow.Price, &productRow.CategoryID)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("error scanning row")
		}
		products = append(products, productRow)
	}

	if rows.Err() != nil {
		log.Println(rows.Err().Error())
		return nil, errors.New("error with rows iteration")
	}

	return products, nil
}
