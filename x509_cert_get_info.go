package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

int x509_certificate_get_info(char *inCert, int inCertLength, int propId, char *outData, int *outDataLength) {
    return kc_funcs->X509CertificateGetInfo(inCert, inCertLength, propId, (unsigned char*)outData, outDataLength);
}
*/
import "C"
import (
	"unsafe"
)

// Obtains field/extension values ​​from a certificate.
//
// The certificate must be previously loaded using one of the following functions:
//
// - LoadKeyStore()
//
// - X509LoadCertificateFromFile()
//
// - X509LoadCertificateFromBuffer().
func (m *Module) X509CertificateGetInfo(inCert string, prop CertProp) (string, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variable and freeing memory when finished.
	cInCert := C.CString(inCert)
	defer C.free(unsafe.Pointer(cInCert))
	outDataLength := 32768
	data := C.malloc(C.ulong(C.sizeof_char * outDataLength))
	defer C.free(data)

	// getting info from certificate.
	rc := int(C.x509_certificate_get_info(cInCert, C.int(len(inCert)), C.int(int(prop)), (*C.char)(data), (*C.int)(unsafe.Pointer(&outDataLength))))

	// return result and checking for errors.
	return C.GoString((*C.char)(data)), m.wrapError(rc)
}
