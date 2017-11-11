package main

/*
#include <stdlib.h>
*/
import "C"

//import "unsafe"

//export HelloWorld
func HelloWorld() *C.char {
	cs := C.CString("Hello World!")
	//defer C.free(unsafe.Pointer(cs))
	return cs
}

func main() {}