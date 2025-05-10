package command

type UpdateUserCommand struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name,omitempty" validate:"lte=100"`
	Email string `json:"email,omitempty" validate:"email"`
}
