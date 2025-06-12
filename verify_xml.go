package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long verify_xml(char *alias, int flags, char *inData, int inDataLength, char *outVerifyInfo, int *outVerifyInfoLen) {
	return kc_funcs->VerifyXML(alias, flags, inData, inDataLength, outVerifyInfo, outVerifyInfoLen);
}
*/
import "C"
import "unsafe"

// Verifies the signature of XML data.
func (m *Module) VerifyXML(alias, inData string, flags Flag) (string, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variables and freeing memory when finished.
	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	cInData := C.CString(inData)
	defer C.free(unsafe.Pointer(cInData))

	outVerifyInfoLen := 64768
	outVerifyInfo := C.malloc(C.ulong(C.sizeof_char * outVerifyInfoLen))
	defer C.free(outVerifyInfo)

	// verifing XML data.
	rc := int(C.verify_xml(cAlias, C.int(flags), cInData, C.int(len(inData)), (*C.char)(outVerifyInfo), (*C.int)(unsafe.Pointer(&outVerifyInfoLen))))
	// checking for errors.
	if err := m.wrapError(rc); err != nil {
		return "", err
	}
	return C.GoString((*C.char)(outVerifyInfo)), nil
}
