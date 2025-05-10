package command

type UpdateUserNameCommand struct {
	ID   string `json:"-" validate:"required"`
	Name string `json:"name,omitempty" validate:"lte=100"`
}
