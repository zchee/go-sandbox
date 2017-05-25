// +build ignore

package main

/*
extern void hoge(int *, int);
inline void foo() {
	int a[] = {0, 1, 2, 3};
	hoge(&a[0], 4);
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	C.foo()
}

//export hoge
func hoge(p *C.int, l C.int) {
	var go_array []C.int
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&go_array))
	slice.Cap = int(l)
	slice.Len = int(l)
	slice.Data = uintptr(unsafe.Pointer(p))
	fmt.Println(go_array)
}
