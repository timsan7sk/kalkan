package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long sign_wsse(char *alias, int flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId) {
    return kc_funcs->SignWSSE(alias, flags, inData, inDataLength, outSign, outSignLength, signNodeId);
}
*/
import "C"
import (
	"encoding/xml"
	"unsafe"
)

const (
	xmlnsSOAP = "http://xmlsoap.org"
	xmlnsWSU  = "http://oasis-open.org"
	replace   = "replace"
)

// Represents soap:Envelope
type soapEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	SOAP    string   `xml:"xmlns:soap,attr"`
	WSU     string   `xml:"xmlns:wsu,attr"`
	Body    soapBody `xml:"soap:Body"`
}

// Represents soap:Body
type soapBody struct {
	ID      string `xml:"wsu:Id,attr"`
	Content string `xml:",chardata"`
}

// Signs SOAP messages according to the WS-Security specification. Required for [SmartBridge]:https://egov.kz.
func (m *Module) SignWSSE(alias, inData, signNodeID string, flags Flag) (string, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variables and freeing memory when finished.
	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	cInData := C.CString(inData)
	defer C.free(unsafe.Pointer(cInData))

	inDataLength := len(inData)
	outSignLength := 50000 + inDataLength
	outSign := C.malloc(C.ulong(C.sizeof_uchar * outSignLength))
	defer C.free(outSign)

	cSignNodeID := C.CString(signNodeID)
	defer C.free(unsafe.Pointer(cSignNodeID))

	// signing data.
	rc := int(C.sign_wsse(cAlias, C.int(flags), cInData, C.int(inDataLength), (*C.uchar)(outSign), (*C.int)(unsafe.Pointer(&outSignLength)), cSignNodeID))

	if err := m.wrapError(rc); err != nil {
		return "", err
	}

	actualLen := int(outSignLength)
	if actualLen <= 0 {
		return "", nil
	}

	signSlice := unsafe.Slice((*byte)(outSign), actualLen)

	if signSlice[actualLen-1] == 0 {
		actualLen--
	}

	return C.GoStringN((*C.char)(outSign), C.int(actualLen)), nil
}
