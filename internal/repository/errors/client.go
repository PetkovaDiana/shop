package errors

import "fmt"

type ErrClientNotFound struct {
	ClientEmail string
}

func (e ErrClientNotFound) Error() string {
	return fmt.Sprintf("client with email: %s not found", e.ClientEmail)
}
