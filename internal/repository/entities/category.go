package entities

const (
	Table_Category        = "category"
	Failed_Category_ID    = "id"
	Failed_Category_title = "title"
)

type Category struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
}
