package helper

import (
	"crypto/sha512"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

func HashGeneratePassword(passwd, salt string) string {
	tempPasswd := pbkdf2.Key([]byte(passwd), []byte(salt), 10000, 256/8, sha512.New)
	// Base64 encoding
	output := base64.StdEncoding.EncodeToString(tempPasswd)
	return output
}

func CreateSalt() string {
	data := make([]byte, 128/8)
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	return sEnc
}
