package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Decryption Program")
	key := []byte("secretToSecretIsGoodPastaAndGood")

	ciphertext, err := ioutil.ReadFile("encrypt.data")
	if err != nil {
		panic(err)
	}

	// Creates cipher block
	cblock, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Creates a Galois Counter Mode
	gcm, err := cipher.NewGCM(cblock)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println("This is not right")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	fmt.Printf("The Nonce: %s \n", string(nonce))
	fmt.Printf("The CipherText: %s \n", string(ciphertext))

	text, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The secret is: %s ", string(text))

}
