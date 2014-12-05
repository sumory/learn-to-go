package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	hash := sha256.New()
	io.WriteString(hash, "just test.")
	fmt.Printf("%x\n", hash.Sum(nil))

	hash = md5.New()
	io.WriteString(hash, "just test.")
	fmt.Printf("%x\n", hash.Sum(nil))
}
