package jwtoken

import "github.com/golang-jwt/jwt"

type Payload struct {
	jwt.StandardClaims
	UserID uint64 `json:"user_id"`
	Email  string `json:"email"`
	Iss    string `json:"iss"`
	Exp    int64  `json:"exp"`
}
