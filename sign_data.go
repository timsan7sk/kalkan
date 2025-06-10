package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <strings.h>
#include "KalkanCrypt.h"

unsigned long sign_data(char *alias, int flag, char *inData, int inDataLength, unsigned char *inSign, int inSignLen, unsigned char *outSign, int *outSignLength) {
     return kc_funcs->SignData(alias, flag, inData, inDataLength, inSign, inSignLen, outSign, outSignLength);
}
*/
import "C"
import "unsafe"

// Signs data.
func (m *Module) SignData(inSign, inData, alias string, flag Flag) (string, error) {
	// lock and ulock mutex.
	m.mu.Lock()
	defer m.mu.Unlock()

	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	cInData := C.CString(inData)
	defer C.free(unsafe.Pointer(cInData))

	inDataLength := len(inData)

	outSignLength := 50000 + 2*inDataLength
	cOutSign := C.malloc(C.ulong(C.sizeof_uchar * outSignLength))
	defer C.free(cOutSign)

	inSignLength := len(inSign)
	cInSign := unsafe.Pointer(C.CString(inSign))
	defer C.free(cInSign)

	rc := int(C.sign_data(cAlias, C.int(int(flag)), cInData, C.int(inDataLength), (*C.uchar)(cInSign), C.int(inSignLength), (*C.uchar)(cOutSign), (*C.int)(unsafe.Pointer(&outSignLength))))

	return C.GoString((*C.char)(cOutSign)), m.wrapError(rc)
}
