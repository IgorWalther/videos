//go:build goexperiment.runtimesecret

package main

import "runtime/secret"

func main() {
	secret.Do(func() {
		// Generate an ephemeral key and
	})

	// func secret_eraseSecrets()
}
