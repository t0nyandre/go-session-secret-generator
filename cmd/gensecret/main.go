package main

import (
	"flag"
	"fmt"

	"gitlab.com/t0nyandre/go-session-secret-generator/pkg/gensecret"
)

func main() {
	hex := flag.Bool("hex", false, "Generate a hex string")
	size := flag.Int("size", 32, "size of secret (in bytes)")
	flag.Parse()

	// No need to go further if we're generating a hex string
	if *hex {
		fmt.Println(gensecret.NewHex(*size))
		return
	}
	fmt.Println(gensecret.NewSecret(*size))
}
