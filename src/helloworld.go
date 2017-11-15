package main

import "C"

//export HelloWorld
func HelloWorld() *C.char {
	cs := C.CString("Hello World!")
	return cs
}

func main() {}