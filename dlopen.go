package kalkan

/*
#include <stdlib.h>
#include <dlfcn.h>
#include <unistd.h>
#include <stdio.h>

#cgo linux LDFLAGS: -ldl
#cgo freebsd LDFLAGS: -ldl
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Obtains pointer to a library (.so).
func dlOpen(libName string) (unsafe.Pointer, error) {
	cLibName := C.CString(libName)
	defer C.free(unsafe.Pointer(cLibName))

	if p := C.dlopen(cLibName, C.RTLD_LAZY); p != nil {
		return p, nil
	}

	return nil, fmt.Errorf("dlOpen: %s", C.GoString(C.dlerror()))
}
