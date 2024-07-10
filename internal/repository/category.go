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

type category struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) Category {
	return &category{db: db}
}

func (c *category) GetAllCategories(ctx context.Context, categoriesIDs []int64) ([]*domainModels.GetAllCategories, error) {
	categories := make([]*domainModels.GetAllCategories, 0, len(categoriesIDs))

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(
			fmt.Sprintf("c.%s, c.%s, COUNT(p.*)",
				entities.Failed_Category_ID,
				entities.Failed_Category_title,
			),
		).
		From(fmt.Sprintf("%s AS c", entities.Table_Category)).
		InnerJoin("product p on c.id = p.category_id").
		GroupBy(fmt.Sprintf("c.%s", entities.Failed_Category_ID)).
		OrderBy(fmt.Sprintf("c.%s", entities.Failed_Category_ID))

	if len(categoriesIDs) != 0 {
		query = query.Where(squirrel.Eq{fmt.Sprintf("c.%s", entities.Failed_Category_ID): categoriesIDs})
	}

	q, p, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := c.db.Query(ctx, q, p...)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("categories not found")
	}
	defer rows.Close()

	for rows.Next() {
		category := new(domainModels.GetAllCategories)
		err = rows.Scan(&category.ID, &category.Name, &category.ProductsCount)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("error scanning row")
		}
		categories = append(categories, category)
	}

	if rows.Err() != nil {
		log.Println(rows.Err().Error())
		return nil, errors.New("error with rows iteration")
	}

	return categories, nil
}
