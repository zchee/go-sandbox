package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: utf8dump <word>")
	}

	w := os.Args[1]
	fmt.Printf("%+q\n", w)
}
