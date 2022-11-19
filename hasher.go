package auth

import "golang.org/x/crypto/bcrypt"

type IPasswordHasher interface {
	Hash(password []byte) ([]byte, error)
	VerifyPassword(correctPassword, password []byte) error
}

type BcryptHasher struct {
	hashedPassword string
}

func NewBCryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (b BcryptHasher) Hash(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (b BcryptHasher) VerifyPassword(correctPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(correctPassword, password)
}

// MockHasher Mock of IPasswordHasher
type MockHasher struct{}

func NewMockHasher() MockHasher {
	return MockHasher{}
}

func (MockHasher) Hash(password []byte) ([]byte, error) {
	return []byte("hashed_password"), nil
}

func (h MockHasher) VerifyPassword(correctPassword, password []byte) error {
	return nil
}
