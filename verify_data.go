package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long verify_data(char *alias, int flags, char *inData, int inDataLength, unsigned char *inoutSign, int inoutSignLength, char *outData, int *outDataLen, char *outVerifyInfo, int *outVerifyInfoLen, int inCertID, char *outCert, int *outCertLength) {
   return kc_funcs->VerifyData(alias, flags, inData, inDataLength, inoutSign, inoutSignLength, outData, outDataLen, outVerifyInfo, outVerifyInfoLen, inCertID, outCert, outCertLength);
}
*/
import "C"
import (
	"unsafe"
)

const (
	// Length of certificate returned from verification.
	outCertLength = 64768
	// length of certificate information returned from verification.
	outVerifyInfoLength = 64768
	// Length of data returned from check.
	outDataLength = 28000
)

// Data returning by VerifyData func.
type VerifiedData struct {
	Cert []byte
	Info []byte
	Data []byte
}

// Verifies signature.
func (m *Module) VerifyData(inSign, inData, alias string, flags Flag) (*VerifiedData, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variables and freeing memory when finished.
	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))
	cInData := C.CString(inData)
	defer C.free(unsafe.Pointer(cInData))
	inDataLength := len(inData)

	cInSign := unsafe.Pointer(C.CString(inSign))
	defer C.free(cInSign)
	inputSignLength := len(inSign)

	var outData [outDataLength]byte
	outDataLen := outDataLength

	var outVerifyInfo [outVerifyInfoLength]byte
	outVerifyInfoLen := outVerifyInfoLength

	inCertID := 0

	var outCert [outCertLength]byte
	outCertLen := outCertLength

	// verifing data.
	rc := int(C.verify_data(cAlias, C.int(flags), cInData, C.int(inDataLength), (*C.uchar)(cInSign), C.int(inputSignLength), (*C.char)(unsafe.Pointer(&outData)),
		(*C.int)(unsafe.Pointer(&outDataLen)), (*C.char)(unsafe.Pointer(&outVerifyInfo)), (*C.int)(unsafe.Pointer(&outVerifyInfoLen)), C.int(inCertID),
		(*C.char)(unsafe.Pointer(&outCert)), (*C.int)(unsafe.Pointer(&outCertLen)),
	))
	// return result anf checking for errors.
	return &VerifiedData{
		Cert: trimSlice(outCert[:]),
		Info: trimSlice(outVerifyInfo[:]),
		Data: trimSlice(outData[:]),
	}, m.wrapError(rc)
}
