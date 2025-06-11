package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

	unsigned long sign_wsse(char *alias, int flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId) {
	    return kc_funcs->SignWSSE(alias, flags, inData, inDataLength, outSign, outSignLength, signNodeId);
	}
*/
import "C"
import "unsafe"

// Signs SOAP messages according to the WS-Security specification. Required for [SmartBridge]:https://sb.egov.kz/.
func (m *Module) SignWSSE(alias, inData, signNodeID string, flags Flag) (string, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variables and freeing memory when finished.
	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	cInData := C.CString(inData)
	defer C.free(unsafe.Pointer(cInData))

	inDataLength := len(inData)
	outSignLength := 50000 + inDataLength
	outSign := C.malloc(C.ulong(C.sizeof_uchar * outSignLength))
	defer C.free(outSign)

	cSignNodeID := C.CString(signNodeID)
	defer C.free(unsafe.Pointer(cSignNodeID))

	// singing data.
	rc := int(C.sign_wsse(cAlias, C.int(flags), cInData, C.int(inDataLength), (*C.uchar)(outSign), (*C.int)(unsafe.Pointer(&outSignLength)), cSignNodeID))

	err := m.wrapError(rc)
	if err != nil {
		return "", err
	}
	return C.GoString((*C.char)(outSign)), nil
}
