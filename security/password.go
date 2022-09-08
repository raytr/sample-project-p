package security

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func Hash(pwd, key string) string {

	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(pwd))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifyPassword(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}
