package password

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	cost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	if err != nil {
		return "", errors.New("FAILED TO PARSE BCRYPT COST")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hashedPassword), nil

}

func ComparePassword(passInput, passFound string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(passFound), []byte(passInput)); err != nil {
		return errors.New("INVALID PASSWORD")
	}

	return nil

}
