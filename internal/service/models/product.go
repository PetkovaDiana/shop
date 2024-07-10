package models

type GetProductsFilter struct {
	ProductsIDs []int64 `json:"products_ids"`
}

type GetAllProducts struct {
	ID          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Price       int64  `json:"price" db:"price"`
	CategoryID  int64  `json:"categories_id" db:"category_id"`
}
