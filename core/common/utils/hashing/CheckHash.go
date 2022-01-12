package hashing

import "golang.org/x/crypto/bcrypt"

func CheckHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
