package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long getLastErrorString(char *errorString, int *bufSize) {
	return kc_funcs->KC_GetLastErrorString(errorString, bufSize);
}
*/
import "C"
import (
	"unsafe"
)

// Length of the error string returned from KalkanCrypt.
const errLength = 65534

// Obtains the text of the last error code and string.
func (m *Module) GetLastErrorString() (ErrorCode, string) {
	errLen := errLength
	var errStr [errLength]byte
	rc := int64(C.getLastErrorString(
		(*C.char)(unsafe.Pointer(&errStr)),
		(*C.int)(unsafe.Pointer(&errLen)),
	))
	return ErrorCode(rc), string(byteSlice(errStr[:]))
}
