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

	m.mu.Lock()
	defer m.mu.Unlock()

	cInCert := C.CString(inCert)
	defer C.free(unsafe.Pointer(cInCert))

	outDataLength := 32768
	data := C.malloc(C.ulong(C.sizeof_char * outDataLength))
	defer C.free(data)

	rc := int(C.x509_certificate_get_info(cInCert, C.int(len(inCert)), C.int(int(prop)), (*C.char)(data), (*C.int)(unsafe.Pointer(&outDataLength))))
	return C.GoString((*C.char)(data)), m.wrapError(rc)
}
