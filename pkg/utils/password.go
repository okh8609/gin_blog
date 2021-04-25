package utils

import (
	"crypto/sha512"
	"encoding/base64"

	"github.com/okh8609/gin_blog/global"
	"golang.org/x/crypto/pbkdf2"
)

func HashPassword(password string) string {
	return base64.StdEncoding.EncodeToString(pbkdf2.Key([]byte(password), []byte(global.Auth.PBKDF2salt), 4096, 128, sha512.New))
}
