package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long get_last_error_string(char *errorString, int *bufSize) {
	return kc_funcs->KC_GetLastErrorString(errorString, bufSize);
}
*/
import "C"
import (
	"unsafe"
)

// Obtains the code and text of the last error.
func (m *Module) GetLastErrorString() (ErrorCode, string) {
	// Length of the error string returned from KalkanCrypt.
	errLen := 65534
	var errStr [65534]byte
	rc := int64(C.get_last_error_string((*C.char)(unsafe.Pointer(&errStr)), (*C.int)(unsafe.Pointer(&errLen))))
	return ErrorCode(rc), string(trimSlice(errStr[:]))
}
