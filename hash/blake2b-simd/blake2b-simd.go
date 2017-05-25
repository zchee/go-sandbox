package main

import (
	"fmt"

	blake2b "github.com/minio/blake2b-simd"
)

func main() {
	s := blake2b.Sum512([]byte("test"))
	fmt.Println(s)
}
