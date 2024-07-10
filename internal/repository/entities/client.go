package entities

const (
	Table_Client           = "client"
	Field_Client_ID        = "id"
	Field_Client_Name      = "name"
	Field_Client_Last_Name = "last_name"
	Field_Client_Number    = "number"
	Field_Client_Password  = "password"
	Field_Client_Email     = "email"
)

type Client struct {
	ID             int64  `db:"id"`
	Name           string `db:"name"`
	LastName       string `db:"last_name"`
	Number         int    `db:"number"`
	PasswordHashed []byte `db:"password_hash"`
	Email          string `db:"email"`
}
