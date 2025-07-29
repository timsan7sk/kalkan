package kalkan

/*
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <dlfcn.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Takes a symbol name and returns a pointer to the symbol or error.
func dlSym(h unsafe.Pointer) (unsafe.Pointer, error) {
	cSym := C.CString("KC_GetFunctionList")
	defer C.free(unsafe.Pointer(cSym))
	C.dlerror()
	p := C.dlsym(h, cSym)

	if err := C.dlerror(); err != nil {
		return nil, fmt.Errorf("dlSym: %s", C.GoString(err))
	}
	return p, nil
}
