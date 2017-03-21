// https://play.golang.org/p/dzO0kdJWJW

package main

import (
	"fmt"
)

func main() {
	for _, i := range []int{1, 2, 23, 26} {
		fmt.Printf("%d %q\n", i, toChar(i))
	}
	for _, i := range []int{1, 2, 23, 26} {
		fmt.Printf("%d %q\n", i, toCharStr(i))
	}
	for _, i := range []int{1, 2, 23, 26} {
		fmt.Printf("%d %q\n", i, toCharStrArr(i))
	}
	for _, i := range []int{1, 2, 23, 26} {
		fmt.Printf("%d %q\n", i, toCharStrConst(i))
	}
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}

func toCharStr(i int) string {
	return string('A' - 1 + i)
}

var arr = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func toCharStrArr(i int) string {
	return arr[i-1]
}

const abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func toCharStrConst(i int) string {
	return abc[i-1 : i]
}
