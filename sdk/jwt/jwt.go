package jwt

import (
	"errors"
	"os"
	u "rakamin-academy/src/entity/user"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user u.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Name,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"id":  user.ID,
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", errors.New("FAILED TO GENERATE TOKEN")
	}

	return signedToken, nil

}

func DecodeToken(token string) (map[string]interface{}, error) {

	decodeToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, errors.New("FAILED TO PARSE TOKEN")
	}

	claims, ok := decodeToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("FAILED TO CLAIMS TOKEN")
	}

	if !decodeToken.Valid {
		return nil, errors.New("INVALID TOKEN")
	}

	return claims, nil

}
