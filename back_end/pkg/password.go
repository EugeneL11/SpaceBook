package pkg

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"

	"golang.org/x/crypto/scrypt"
)

const (
	saltSize = 16
	hashSize = 16
	N        = 16384
	r        = 8
	p        = 1
)

// GeneratePasswordHash generates a 128-bit password hash using scrypt.
func GeneratePasswordHash(password string) (*big.Int, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	hash, err := scrypt.Key([]byte(password), salt, N, r, p, hashSize)
	if err != nil {
		return nil, err
	}

	hashInt := new(big.Int).SetBytes(hash)
	return hashInt, nil
}

// StoreHashInDatabase stores the password hash in the database.
func StoreHashInDatabase(db *sql.DB, hash *big.Int) error {
	// Convert the hash to a string for storage in the database
	hashString := hash.String()

	// Store the hash in the database (replace "users" and "password_hash" with your table and column names)
	_, err := db.Exec("INSERT INTO users (password_hash) VALUES ($1)", hashString)
	return err
}

// VerifyPassword verifies a password against the stored hash in the database.
func VerifyPassword(db *sql.DB, password string) error {
	// Retrieve the stored hash from the database (replace "users" and "password_hash" with your table and column names)
	var storedHash string
	err := db.QueryRow("SELECT password_hash FROM users WHERE username = $1", "desired_username").Scan(&storedHash)
	if err != nil {
		return err
	}

	// Convert the stored hash string to a big integer
	storedHashInt, success := new(big.Int).SetString(storedHash, 10)
	if !success {
		return fmt.Errorf("failed to convert stored hash to big.Int")
	}

	// Generate a hash from the provided password
	newHash, err := GeneratePasswordHash(password)
	if err != nil {
		return err
	}

	// Compare the generated hash with the stored hash
	if newHash.Cmp(storedHashInt) != 0 {
		return fmt.Errorf("password does not match")
	}

	fmt.Println("Password is correct")
	return nil
}
