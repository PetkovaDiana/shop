package entities

const (
	Table_Product              = "product"
	Failed_Product_ID          = "id"
	Failed_Product_Title       = "title"
	Failed_Product_Description = "description"
	Failed_Product_Price       = "price"
	Faile_Category_id          = "category_id"
)

type Product struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
	CategoryID  string `db:"category_id"`
}
