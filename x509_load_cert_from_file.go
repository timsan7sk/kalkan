package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

int x509_load_certificate_from_file(char *certPath, int certType) {
    return kc_funcs->X509LoadCertificateFromFile(certPath, certType);
}
*/
import "C"
import (
	"unsafe"
)

// Loads cerificate from file.
func (m *Module) X509LoadCertificateFromFile(certPath string, certType CertType) error {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variable and freeing memory when finished.
	cCertPath := C.CString(certPath)
	defer C.free(unsafe.Pointer(cCertPath))

	// loading certificate.
	rc := int(C.x509_load_certificate_from_file(cCertPath, C.int(int(certType))))

	// checking for errors.
	return m.wrapError(rc)
}
