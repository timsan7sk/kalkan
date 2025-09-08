package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <string.h>
#include "KalkanCrypt.h"

unsigned long hash_data(char *algorithm, int flags, char *inData, int inDataLength, unsigned char *outData, int *outDataLength) {
	return kc_funcs->HashData(algorithm, flags, inData, inDataLength, outData, outDataLength);
 }
*/
import "C"
import "unsafe"

// Hasing algorithm.
type HashAlg string

const (
	HashAlgSHA256 HashAlg = "sha256"
	HashAlgGOST95 HashAlg = "Gost34311_95"
	HashAlgGOST15 HashAlg = "Gost34311_15"
)

// Hashes data.
func (m *Module) HashData(alg HashAlg, flags Flag, inData string) (string, error) {

	m.mu.Lock()
	defer m.mu.Unlock()
	cAlg := C.CString(string(alg))
	defer C.free(unsafe.Pointer(cAlg))

	cInData := C.CString(inData)
	defer C.free(unsafe.Pointer(cInData))
	inDataLength := len(inData)

	outDataLength := 2000
	outData := C.malloc(C.ulong(C.sizeof_uchar * outDataLength))
	defer C.free(outData)
	C.memset(outData, 0, C.ulong(C.sizeof_uchar*outDataLength))

	rc := int(C.hash_data(cAlg, C.int(int(flags)), cInData, C.int(inDataLength), (*C.uchar)(outData), (*C.int)(unsafe.Pointer(&outDataLength))))

	return C.GoString((*C.char)(outData)), m.wrapError(rc)

}
