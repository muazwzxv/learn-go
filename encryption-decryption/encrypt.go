package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	fmt.Println("Encryption Program")

	text := []byte("My secret to be encrypt, I love hardcore and dominate sex (This is an ester egg, hmu if someone see this hehe)")
	key := []byte("secretToSecretIsGoodPastaAndGood") // Must be 32 bytes

	// Creates a new cipher
	// func NewCipher(key []byte) (cipher.Block, error)
	cblock, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// gcm = Galois Counter Mode
	gcm, err := cipher.NewGCM(cblock)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// Seal the data
	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	encrypted := gcm.Seal(nonce, nonce, text, nil)
	fmt.Println(encrypted)
}
