package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	s := "sha256 this string"

	h := sha256.New()
	h1 := sha512.New()

	h.Write([]byte(s))
	bs := h.Sum(nil)

	h1.Write([]byte(s))
	bs1 := h1.Sum(nil)

	fmt.Println(s)
	fmt.Printf("sha256: %x\n", bs)
	fmt.Printf("sha512: %x\n", bs1)
}
