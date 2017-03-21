// https://gist.github.com/jedy/4167513

package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

func main() {
	k := struct {
		Type   uint32
		Whence uint32
		Start  uint64
		Len    uint64
		Pid    uint32
	}{
		syscall.F_WRLCK,
		uint32(os.SEEK_SET),
		0,
		10,
		0,
	}

	f, err := os.OpenFile("./x", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	y, z, err := syscall.Syscall(syscall.SYS_FCNTL, f.Fd(), uintptr(syscall.F_SETLKW), uintptr(unsafe.Pointer(&k)))
	fmt.Println(y)
	fmt.Println(z)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
