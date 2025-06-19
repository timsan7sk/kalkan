package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long get_cert_from_xml(char *inXML, int inXMLLen, int inSignId, char *outCert, int *outCertLength) {
    return kc_funcs->KC_getCertFromXML(inXML, inXMLLen, inSignId, outCert, outCertLength);
}
*/
import "C"
import "unsafe"

// Obtains a certificate from a data signature in XML format.
func (m *Module) GetCertFromXML(xml string, signID int) (string, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()
	// preparing variable and freeing memory when finished.
	cXML := C.CString(xml)
	defer C.free(unsafe.Pointer(cXML))
	outCertLen := 32768
	outCert := C.malloc(C.ulong(C.sizeof_uchar * outCertLen))
	defer C.free(outCert)

	rc := int(C.get_cert_from_xml(cXML, C.int(len(xml)), C.int(signID), (*C.char)(outCert), (*C.int)(unsafe.Pointer(&outCertLen))))

	// checking for errors.
	if err := m.wrapError(rc); err != nil {
		return "", err
	}
	return C.GoString((*C.char)(outCert)), nil
}
