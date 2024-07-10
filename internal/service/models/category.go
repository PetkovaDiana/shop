package models

type GetCategoriesFilter struct {
	CategoriesIDs []int64 `json:"categories_ids"`
}

type GetAllCategories struct {
	ID            int64  `json:"id" db:"id"`
	Name          string `json:"name" db:"title"`
	ProductsCount int64  `json:"products_count" db:"products_count"`
}
