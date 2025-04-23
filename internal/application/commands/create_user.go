package commands

type CreateUserCommand struct {
	Name           string `json:"name"`
	DocumentNumber string `json:"document_number"`
	DocumentType   string `json:"document_type"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}
