package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long get_sig_alg_from_xml(const char* xml_in, int xml_in_size, char *retSigAlg, int *retLen) {
	return kc_funcs->KC_getSigAlgFromXML(xml_in, xml_in_size, retSigAlg, retLen);
}
*/
import "C"
import "unsafe"

// Obtains the signature algorithm from XML.
func (m *Module) GetSignAlgorithmFromXML(xml string) (string, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()
	// preparing variable and freeing memory when finished.
	cXML := C.CString(xml)
	defer C.free(unsafe.Pointer(cXML))
	cAlg := C.malloc(C.ulong(C.sizeof_uchar * 1024))
	defer C.free(cAlg)
	algLen := 1024
	rc := int(C.get_sig_alg_from_xml(cXML, C.int(len(xml)), (*C.char)(cAlg), (*C.int)(unsafe.Pointer(&algLen))))

	return C.GoString((*C.char)(cAlg)), m.wrapError(rc)
}
