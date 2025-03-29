package common

import "golang.org/x/crypto/bcrypt"

const (
	HASH_COST = 12
)

func EncryptPassword(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), HASH_COST)
	if err != nil {
		return nil, err
	}
	return hashed, nil
}
