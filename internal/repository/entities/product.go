package entities

const (
	Table_Product              = "product"
	Failed_Product_ID          = "id"
	Failed_Product_title       = "title"
	Failed_Product_description = "description"
	Failed_Product_price       = "price"
)

type Product struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
	CategoryID  string `db:"category_id"`
}
