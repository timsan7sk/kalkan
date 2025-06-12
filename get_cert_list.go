package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long get_certificates_list(char *certificates, unsigned long *cert_count) {
    return kc_funcs->KC_GetCertificatesList(certificates, cert_count);
}
*/
import "C"
import "unsafe"

// Obtains a list of certificates as a string and their number.
func (m *Module) GetCertList() (string, int, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()
	// preparing variable and freeing memory when finished.
	var count int
	certsLen := 4096
	cCerts := C.malloc(C.ulong(C.sizeof_char * certsLen))
	defer C.free(cCerts)

	rc := int(C.get_certificates_list((*C.char)(cCerts), (*C.ulong)(unsafe.Pointer(&count))))

	err := m.wrapError(rc)
	if err != nil {
		return "", 0, err
	}
	return C.GoString((*C.char)(cCerts)), count, nil
}
