package password

import "golang.org/x/crypto/bcrypt"

type PasswordComparer interface {
	ComparePassword(hashedPassword, plainPassword string) error
	HashPassword(password string) (string, error)
}

// BcryptPasswordComparer implements the PasswordComparer interface using bcrypt.
type BcryptPasswordComparer struct{}

// NewPasswordComparer creates a new instance of PasswordComparer using bcrypt.
func NewPasswordComparer() PasswordComparer {
	return &BcryptPasswordComparer{}
}

// HashPassword hashes a plain password using bcrypt.
func (c *BcryptPasswordComparer) HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// ComparePassword compares a hashed password with a plain password using bcrypt.
func (c *BcryptPasswordComparer) ComparePassword(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return err
	}
	return nil
}
