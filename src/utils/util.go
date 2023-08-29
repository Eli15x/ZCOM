package utils

import (
	"github.com/satori/go.uuid"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func CreateCodeId() string {
	return uuid.NewV4().String()
}

func Encrypt(code string) string{
	key := []byte("ZCOM2023")
	plaintext := []byte(code)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte("ZCO")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	ciphertextStr := string(ciphertext[:])
	return ciphertextStr
}

func decrypt(code string) string{
	key := []byte("ZCOM2023")
	ciphertext, _ := hex.DecodeString(code)
	nonce := []byte("ZCO")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Plaintext: %s\n", string(plaintext))
	plaintextStr := string(plaintext[:])
	return plaintextStr
}


