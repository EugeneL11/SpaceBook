package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12) // Error can only occur for password greater than 72 bytes
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func VerifyPassword(password string, hashedPassword []byte) bool {
	// Compare the hashed password with the provided password
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}
