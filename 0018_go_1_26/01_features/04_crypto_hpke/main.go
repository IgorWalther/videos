package main

import (
	"crypto/hpke"
	"fmt"
)

// Encrypt a single message from a sender to a recipient using the one-shot API.
var kem, kdf, aead = hpke.MLKEM768X25519(), hpke.HKDFSHA256(), hpke.AES256GCM()

var (
	recipientPrivateKey hpke.PrivateKey
	publicKeyBytes      []byte
	ciphertext          []byte
)

func init() {
	k, err := kem.GenerateKey()
	if err != nil {
		panic(err)
	}
	recipientPrivateKey = k
	publicKeyBytes = k.PublicKey().Bytes()
}

func server() {
	plaintext, err := hpke.Open(recipientPrivateKey, kdf, aead, []byte("public"), ciphertext)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted: %s\n", plaintext)
}

func client() {
	publicKey, err := kem.NewPublicKey(publicKeyBytes)

	if err != nil {
		panic(err)
	}

	message := []byte("secret message")

	ct, err := hpke.Seal(publicKey, kdf, aead, []byte("public"), message)

	if err != nil {
		panic(err)
	}

	ciphertext = ct
}

func main() {
	client()
	server()
}
