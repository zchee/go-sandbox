// https://gotipplay.golang.org/p/DS5wA_NoBPv

package main

import (
	"fmt"
	"testing"
	"unsafe"
)

// combineed _type and maptype structs.
// - https://cs.opensource.google/go/go/+/refs/tags/go1.18.1:src/runtime/type.go;l=34-50
// - https://cs.opensource.google/go/go/+/refs/tags/go1.18.1:src/runtime/type.go;l=348-354
type mh struct {
	// type _type
	_ uintptr                                   // size
	_ uintptr                                   // ptrdata
	_ uint32                                    // hash
	_ uint8                                     // tflag
	_ uint8                                     // align
	_ uint8                                     // fieldAsign
	_ uint8                                     // kind
	_ func(unsafe.Pointer, unsafe.Pointer) bool // equal
	_ *byte                                     // gcdata
	_ int32                                     // str
	_ int32                                     // ptrToThis
	// type maptype
	_ unsafe.Pointer // key
	_ unsafe.Pointer // elem
	_ unsafe.Pointer // bucket

	// function for hashing keys (ptr to key, seed) -> hash
	hasher func(unsafe.Pointer, uintptr) uintptr
}

func Hash[T comparable](t T) uintptr {
	var m interface{} = map[T]struct{}(nil)
	hasher := (*mh)(*(*unsafe.Pointer)(unsafe.Pointer(&m))).hasher

	return hasher(noescape(unsafe.Pointer(&t)), 0)
}

//go:noescape
//go:nosplit
//go:nocheckptr
//go:linkname noescape runtime.noescape
func noescape(p unsafe.Pointer) unsafe.Pointer

func main() {
	var sum uintptr

	fmt.Println(testing.AllocsPerRun(1000, func() {
		sum += Hash(0)
	}))

	fmt.Println(sum)
}
