package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long sign_xml(char *alias, int flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId, char *parentSignNode, char *parentNameSpace) {
	return kc_funcs->SignXML(alias, flags, inData, inDataLength, outSign, outSignLength, signNodeId, parentSignNode, parentNameSpace);
}
*/
import "C"
import (
	"unsafe"
)

// Signs data in XML format.
func (m *Module) SignXML(xml, alias string, flags Flag, signNodeID, parentSignNode, parentNameSpace string) (signedXML string, err error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variables and freeing memory when finished.
	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	cInData := C.CString(xml)
	defer C.free(unsafe.Pointer(cInData))

	inDataLength := len(xml)
	outSignLength := 50000 + inDataLength
	outSign := C.malloc(C.ulong(C.sizeof_uchar * outSignLength))
	defer C.free(outSign)

	cSignNodeID := C.CString(signNodeID)
	defer C.free(unsafe.Pointer(cSignNodeID))

	cParentSignNode := C.CString(parentSignNode)
	defer C.free(unsafe.Pointer(cParentSignNode))

	cParentNameSpace := C.CString(parentNameSpace)
	defer C.free(unsafe.Pointer(cParentNameSpace))

	// singing XML data.
	rc := int(C.sign_xml(cAlias, C.int(flags), cInData, C.int(inDataLength), (*C.uchar)(outSign), (*C.int)(unsafe.Pointer(&outSignLength)), cSignNodeID, cParentSignNode, cParentNameSpace))
	// return result and checking for error.
	return C.GoString((*C.char)(outSign)), m.wrapError(rc)
}
