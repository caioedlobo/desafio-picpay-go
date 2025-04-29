package persistence

import (
	"context"
	"database/sql"
	"errors"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *user.User) error {
	query := `
        INSERT INTO users (name, document_number, document_type, email, password_hash, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.DocumentNumber,
		user.DocumentType,
		user.Email,
		user.Password.GetHash(),
		user.CreatedAt,
	).Scan(&user.ID)

	return err
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id int64) (*user.User, error) {
	query := `
        SELECT id, name, document_number, document_type, email, password_hash, created_at
        FROM users
        WHERE id = $1
    `

	var user user.User
	var passwordHash []byte

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.DocumentNumber,
		&user.DocumentType,
		&user.Email,
		&passwordHash,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, nil
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `
        SELECT id, name, document_number, document_type, email, password_hash, created_at
        FROM users
        WHERE email = $1
    `

	var user user.User
	var passwordHash []byte

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.DocumentNumber,
		&user.DocumentType,
		&user.Email,
		&passwordHash,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	//user.Password = model.Password{hash: passwordHash}

	return &user, nil
}

func (r *PostgresUserRepository) FindByDocument(ctx context.Context, documentNumber string, documentType user.DocumentType) (*user.User, error) {
	query := `
        SELECT id, name, document_number, document_type, email, password_hash, created_at
        FROM users
        WHERE document_number = $1 AND document_type = $2
    `

	var user user.User
	var passwordHash []byte

	err := r.db.QueryRowContext(ctx, query, documentNumber, documentType).Scan(
		&user.ID,
		&user.Name,
		&user.DocumentNumber,
		&user.DocumentType,
		&user.Email,
		&passwordHash,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	//user.Password = model.Password{hash: passwordHash}

	return &user, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *user.User) error {
	query := `
        UPDATE users
        SET name = $1, email = $2
        WHERE id = $3
    `

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.ID,
	)

	return err
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id int64) error {
	query := `
        DELETE FROM users
        WHERE id = $1
    `

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
