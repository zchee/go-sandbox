package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: utf8dump word [word...]")
	}

	for _, s := range os.Args[1:] {
		if idx := strings.Index(s, "U"); idx > -1 {
			s = s[idx+2:]
			u, _ := strconv.ParseUint(s, 16, 32)
			r := rune(u)
			fmt.Printf("%2s\n%2s\n", s, string(r))
		}
		fmt.Printf("%2s\n%+2q\n", s, s)
	}
}
