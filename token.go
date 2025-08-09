package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

var secretKey = []byte("OzBaClaveUltraSecretaDe32Bytes!!") // Debe tener 32 bytes

func GenerateToken(uuid, mac string) string {
	data := uuid + mac + string(secretKey)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func EncryptAndSaveToken(token string, filename string) error {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return err
	}

	plaintext := []byte(token)
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, secretKey[:aes.BlockSize])
	stream.XORKeyStream(ciphertext, plaintext)
	fmt.Println("Token cifrado exitosamente:", ciphertext)
	return ioutil.WriteFile(filename, ciphertext, 0644)
	//fmt.Println("Token guardado exitosamente en", filename);
}
