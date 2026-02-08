package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func main() {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	fmt.Println(key.D)

	prim, _ := rand.Prime(rand.Reader, 64)
	fmt.Println(prim)
}
