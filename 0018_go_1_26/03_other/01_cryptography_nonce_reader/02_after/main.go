package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func main() {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), nil)
	fmt.Println(key.D)

	prim, _ := rand.Prime(nil, 64)
	fmt.Println(prim)
}
