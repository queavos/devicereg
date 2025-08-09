package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
)

// Usa la misma clave que en token.go
//var secretKey = []byte("OzBaClaveUltraSecretaDe32Bytes!!") // 32 bytes exactos

func LoadAndDecryptToken(filename string) (string, error) {
	ciphertext, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, secretKey[:aes.BlockSize])
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
