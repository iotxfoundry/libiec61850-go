package libiec61850go

/*
#include <stdint.h>

*/
import "C"
import "unsafe"

func StringData(str string) *C.char {
	return (*C.char)(unsafe.Pointer(unsafe.StringData(str + "\x00")))
}

func SliceData(str []byte) *C.uint8_t {
	return (*C.uint8_t)(unsafe.Pointer(unsafe.SliceData(str)))
}
