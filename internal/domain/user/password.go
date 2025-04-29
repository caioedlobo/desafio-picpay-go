package user

import "golang.org/x/crypto/bcrypt"

type Password struct {
	plaintext *string
	hash      []byte
}

func NewPassword(plaintext string) (*Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return nil, err
	}
	return &Password{plaintext: &plaintext, hash: hash}, nil
}

func (p *Password) Set(plaintext string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return []byte(""), err
	}
	return hash, nil
}

func (p *Password) GetHash() []byte {
	return p.hash
}
