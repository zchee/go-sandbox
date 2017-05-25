package main

import "fmt"

type Test int

const (
	iota1 Test = 1 << iota
	iota2
	iota3
	iota4
	iota5
)

func main() {
	fmt.Println(iota1)
	fmt.Println(iota2)
	fmt.Println(iota3)
	fmt.Println(iota4)
	fmt.Println(iota5)
}
