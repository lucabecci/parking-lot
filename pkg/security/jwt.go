package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lucabecci/parking-lot/pkg"
)

const (
	JWT_SECRET   = "secret"
	JWT_EXP_HOUR = 1
	JWT_EXP_MIN  = 0
	JWT_EXP_SEC  = 40
)

func NewToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"nbf": time.Now().Unix(),
		"lat": time.Now().Unix(),
		"exp": time.Now().Local().Add(time.Hour*time.Duration(JWT_EXP_HOUR) + time.Minute*time.Duration(JWT_EXP_MIN) + time.Second*time.Duration(JWT_EXP_SEC)).Unix(),
	})

	sToken, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", pkg.ErrToCreateToken
	}
	return sToken, nil
}
