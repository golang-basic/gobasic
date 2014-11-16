// Created by cgo - DO NOT EDIT

//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:1
package cgoexamples
//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:20

//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:19
import "unsafe"
//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:22

//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:21
const salt string = "$6$16814ad1efd0d232$"
//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:24

//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:23
type Crypt interface {
	EncryptWithSalt(key string) string
}
//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:28

//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:27
type UnixCrypt struct{}
//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:30

//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:29
func NewUnixCrypt() Crypt {
	return UnixCrypt{}
}
//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:34

//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:33
func (uc UnixCrypt) EncryptWithSalt(key string) string {
	ckey := _Cfunc_CString(key)
	csalt := _Cfunc_CString(salt)
	out := _Cfunc_GoString(_Cfunc_crypt(ckey, csalt))
	_Cfunc_free(unsafe.Pointer(ckey))
	_Cfunc_free(unsafe.Pointer(csalt))
	return out
}
//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:43

//line /Users/sonis/gobasic/basics/cgoexamples/cgoexample.go:42
func Example() {
	cs := _Cfunc_CString("Hello from stdio\n")
	_Cfunc_myprint(cs)
	_Cfunc_free(unsafe.Pointer(cs))
}
