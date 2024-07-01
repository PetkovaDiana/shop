package repository

import (
	"errors"
	"fmt"
	domainModels "github.com/PetkovaDiana/shop/internal/service/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type category struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) Category {
	return &category{db: db}
}

func (c *category) GetAllCategories() ([]domainModels.GetAllCategories, error) {
	var categories []domainModels.GetAllCategories

	query := fmt.Sprintf(`
		SELECT 
		    c.id, 
		    c.title, 
		    COUNT(p.*) AS products_count 
		FROM category c
    		INNER JOIN product p on c.id = p.category_id
		GROUP BY c.id 
		ORDER BY c.id;`)

	err := c.db.Select(&categories, query)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("not a found")
	}

	return categories, nil
}
