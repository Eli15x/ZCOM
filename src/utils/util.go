package utils

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
)

func CreateCodeId() string {
	return uuid.NewV4().String()
}

func Encrypt(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

// Função para criar um ZIP de todos os XMLs em um diretório
func CreateZipFromDirectory(directory string) ([]byte, error) {
	var buffer bytes.Buffer
	zipWriter := zip.NewWriter(&buffer)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ".xml" {
			return nil
		}

		// Adiciona o arquivo XML ao ZIP
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		zipFile, err := zipWriter.Create(info.Name())
		if err != nil {
			return err
		}

		_, err = io.Copy(zipFile, file)
		return err
	})

	if err != nil {
		return nil, err
	}

	// Fecha o zipWriter para finalizar o arquivo ZIP
	if err := zipWriter.Close(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
