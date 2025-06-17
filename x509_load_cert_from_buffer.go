package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

int x509_load_certificate_from_buffer(char *inCert, int certLength, int flag) {
    return kc_funcs->X509LoadCertificateFromBuffer((unsigned char*)inCert, certLength, flag);
}
*/
import "C"
import (
	"unsafe"
)

// Loads a certificate from a byte slice.
func (m *Module) X509LoadCertificateFromBuffer(inCert []byte, flags CertEncodingType) error {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()
	// loading certificate.
	rc := int(C.x509_load_certificate_from_buffer((*C.char)(unsafe.Pointer(&inCert)), C.int(len(inCert)), C.int(int(flags))))
	// checking for errors.
	return m.wrapError(rc)
}
