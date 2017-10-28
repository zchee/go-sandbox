// Command crypto-rand generate random base64 string using crypto/rand.
//  https://stackoverflow.com/q/32349807/5228839
//
// This snippet for:
//  https://cloud.google.com/storage/docs/using-encryption-keys?hl=ja
package main

import (
	"fmt"
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	fmt.Println(string(b))
	return base64.URLEncoding.EncodeToString(b), err
}

func main() {
	// Example: this will give us a 44 byte, base64 encoded output
	token, err := GenerateRandomString(32)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
	}
	fmt.Println(token)
}
