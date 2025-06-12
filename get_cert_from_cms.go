package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

long get_cert_from_cms(char *inCMS, int inCMSLen, int inSignId, int flags, char *outCert, int *outCertLength) {
	return kc_funcs->KC_GetCertFromCMS(inCMS, inCMSLen, inSignId, flags, outCert, outCertLength);
}
*/
import "C"
import "unsafe"

// Obtains a certificate from a signature in CMS format.
func (m *Module) GetCertFromCMS(cms string, signID int, flag Flag) (string, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variable and freeing memory when finished.
	cCMS := C.CString(cms)
	defer C.free(unsafe.Pointer(cCMS))
	outCertLen := 32768
	outCert := C.malloc(C.ulong(C.sizeof_uchar * outCertLen))
	defer C.free(outCert)

	// gitting certificate from cms.
	rc := int(C.get_cert_from_cms(cCMS, C.int(len(cms)), C.int(signID), C.int(int(flag)), (*C.char)(outCert), (*C.int)(unsafe.Pointer(&outCertLen))))
	// checking for errors.
	if err := m.wrapError(rc); err != nil {
		return "", err
	}
	return C.GoString((*C.char)(outCert)), nil
}
