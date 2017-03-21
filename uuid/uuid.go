package main

/*
#include <stdio.h>
#include <uuid/uuid.h>
char uuid_str[37];
extern inline char* uuidgen() {
	uuid_t uuid;
	uuid_generate_random(uuid);
	uuid_unparse_upper(uuid, uuid_str);
	return uuid_str;
}
*/
import "C"
import "fmt"

func uuidgen() string {
	return C.GoString(C.uuidgen())
}

func main() {
	fmt.Println(uuidgen())
}
