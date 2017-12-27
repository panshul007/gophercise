package rand

import (
	"crypto/rand"
	"encoding/base64"
)

const RemeberTokenBytes = 32

// Bytes will help us generate n random bytes.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// String will generate a byte slice of size nBytes and then return a string
// that is the base64 encoded string of generated bytes.
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RememberToken is a helper function to generate remember token of predetermined size.
func RemeberToken() (string, error) {
	return String(RemeberTokenBytes)
}
