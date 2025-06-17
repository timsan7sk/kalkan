package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <strings.h>
// #include "KalkanCrypt.h"
//
// int x509_validate_certificate(char *inCert, int inCertLength, int validType, char *validPath, long long checkTime, char *outInfo, int *outInfoLength, int flags, char* getResp, int *getRespLength) {
//     int rc = kc_funcs->X509ValidateCertificate(inCert, inCertLength, validType, validPath, checkTime, outInfo, outInfoLength, flags, getResp, getRespLength);
//     for(size_t i = 0; i < *outInfoLength-1; i++)
//         if(outInfo[i] == '\r' && outInfo[i+1] != '\n') outInfo[i] = '\n';
//     return rc;
// }
import "C"
import (
	"unsafe"
)

const (
	length = 8192
)

// Validates the certificate.
func (m *Module) X509ValidateCertificate(inCert string, validateType ValidateType, validatePath string) (string, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variable and freeing memory when finished.
	cInCert := C.CString(inCert)
	defer C.free(unsafe.Pointer(cInCert))

	cValidPath := C.CString(validatePath)
	defer C.free(unsafe.Pointer(cValidPath))

	var kcOutInfo [length]byte
	kcOutInfoLen := length

	var kcGetResp [length]byte
	kcGetRespLen := length

	// validating certificate.
	rc := int(C.x509_validate_certificate(cInCert, C.int(len(inCert)), C.int(int(validateType)), cValidPath, 0, (*C.char)(unsafe.Pointer(&kcOutInfo)), (*C.int)(unsafe.Pointer(&kcOutInfoLen)), 0, (*C.char)(unsafe.Pointer(&kcGetResp)), (*C.int)(unsafe.Pointer(&kcGetRespLen))))

	// return result and checking for errors.
	return string(trimSlice(kcOutInfo[:])), m.wrapError(rc)
}
