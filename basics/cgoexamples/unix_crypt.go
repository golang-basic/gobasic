package cgoexamples

// #cgo linux LDFLAGS: -lcrypt
// #define _GNU_SOURCE
// #if defined(__APPLE__)
// #include <unistd.h>
// #else
// #include <crypt.h>
// #endif
// #include <stdlib.h>
import "C"

import "unsafe"

const salt string = "$6$16814ad1efd0d232$"

type Crypt interface {
	EncryptWithSalt(key string) string
}

type UnixCrypt struct{}

func NewUnixCrypt() Crypt {
	return UnixCrypt{}
}

func (uc UnixCrypt) EncryptWithSalt(key string) string {
	ckey := C.CString(key)
	csalt := C.CString(salt)
	out := C.GoString(C.crypt(ckey, csalt))
	C.free(unsafe.Pointer(ckey))
	C.free(unsafe.Pointer(csalt))
	return out
}
