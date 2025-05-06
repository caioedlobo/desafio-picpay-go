package command

type CreateUserCommand struct {
	Name           string `json:"name" validate:"required"`
	DocumentNumber string `json:"document_number" validate:"required,gte=11,lte=14"`
	DocumentType   string `json:"document_type" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required"`
}
