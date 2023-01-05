package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	size := flag.Int("size", 0, "size of session secret")
	flag.Parse()

	if *size == 0 {
		*size = 16
	}

	b := make([]byte, *size)
	_, err := rand.Read(b)
	if err != nil {
		panic(fmt.Sprintf("Error creating session secret: %s", err))
	}

	s := base64.StdEncoding.EncodeToString(b)
	fmt.Println(s)
}
