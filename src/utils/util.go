package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/satori/go.uuid"
                                                                                                                                                                                     


)

func CreateCodeId() string {
	return uuid.NewV4().String()
}

func Encrypt(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}
