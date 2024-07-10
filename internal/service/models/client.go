package models

type Client struct {
	ID int64 `json:"id" db:"id"`
	CreateClient
}

type CreateClient struct {
	Name           string `json:"name" db:"name"`
	LastName       string `json:"last_name" db:"last_name"`
	Number         int    `json:"number" db:"number"`
	Password       string `json:"password,omitempty"`
	PasswordHashed []byte `json:"-" db:"password_hash"`
	Email          string `json:"email" db:"email"`
}

type AuthClient struct {
	Email          string `json:"email" db:"email"`
	Password       string `json:"password,omitempty"`
	PasswordHashed []byte `json:"-" db:"password_hash"`
}
