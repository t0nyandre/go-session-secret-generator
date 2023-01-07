package gensecret

import (
	"crypto/rand"
	"fmt"
	rand_math "math/rand"
	"time"
)

// Characters to use for generating a secret
const Charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-+?"

func generateBytes(size int) []byte {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		panic(fmt.Sprintf("Error generating secret: %s", err))
	}
	return b
}

func generateSecret(size int) []byte {
	rand_math.Seed(time.Now().UnixNano())
	secret := generateBytes(size)
	for i := range secret {
		secret[i] = Charset[rand_math.Intn(len(Charset))]
	}
	return secret
}

func NewHex(size int) string {
	return fmt.Sprintf("%x", generateBytes(size))
}

func NewSecret(size int) string {
	return fmt.Sprintf("%s", generateSecret(size))
}
