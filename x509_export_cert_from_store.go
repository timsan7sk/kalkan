package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

int X509_export_certificate_from_store(char *alias, int flag, char *outCert, int *outCertLength) {
    return kc_funcs->X509ExportCertificateFromStore(alias, flag, outCert, outCertLength);
}
*/
import "C"
import (
	"unsafe"
)

// Exports a certificate from the store.
func (m *Module) X509ExportCertificateFromStore(alias string) (string, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	flags := 0
	outCertLen := 32768

	// preparing variable and freeing memory when finished.
	cert := C.malloc(C.ulong(C.sizeof_char * outCertLen))
	defer C.free(cert)
	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	// exporting certificate fro the store.
	rc := int(C.X509_export_certificate_from_store(cAlias, C.int(flags), (*C.char)(cert), (*C.int)(unsafe.Pointer(&outCertLen))))

	// returning result and checking for errors.
	return C.GoString((*C.char)(cert)), m.wrapError(rc)
}
