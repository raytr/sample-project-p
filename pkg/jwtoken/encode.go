package jwtoken

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtokenEncode struct {
	secret string
	iss    string
	exp    time.Duration
}

func NewJwtokenEncode(secret, iss string, exp time.Duration) *JwtokenEncode {
	return &JwtokenEncode{
		secret: secret,
		iss:    iss,
		exp:    exp,
	}
}

func (sa *JwtokenEncode) Encode(userID uint64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Payload{
		UserID: userID,
		Email:  email,
		Exp:    time.Now().Add(sa.exp).Unix(),
		Iss:    sa.iss,
	})
	tokenString, err := token.SigningString(sa.secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
