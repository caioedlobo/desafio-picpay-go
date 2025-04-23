package queries

type GetUserByIDQuery struct {
	ID int64 `json:"id"`
}

type GetUserByEmailQuery struct {
	Email string `json:"email"`
}

type GetUserByDocumentQuery struct {
	DocumentNumber string `json:"document_number"`
	DocumentType   string `json:"document_type"`
}
