package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func CheckPassword(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
