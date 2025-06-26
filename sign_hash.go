package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <strings.h>
#include "KalkanCrypt.h"

	unsigned long sign_hash(char *alias, int flags, char *inHash, int inHashLength, unsigned char *outSign, int *outSignLength) {
	     return kc_funcs->SignHash(alias, flags, inHash, inHashLength, outSign, outSignLength);
	}
*/
import "C"
import "unsafe"

// Signs the input hashed data.
func (m *Module) SignHash(alias, inHash string, flags Flag) (string, error) {
	// lock and ulock mutex.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variables and freeing memory when finished.
	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	cInHash := C.CString(inHash)
	defer C.free(unsafe.Pointer(&cInHash))
	inHashLength := len(inHash)
	outSignLength := 50000 + 2*inHashLength
	cOutSign := C.malloc(C.ulong(C.sizeof_uchar * outSignLength))
	defer C.free(cOutSign)
	// singing hashed data.
	rc := int(C.sign_hash(cAlias, C.int(int(flags)), cInHash, C.int(inHashLength), (*C.uchar)(cOutSign), (*C.int)(unsafe.Pointer(&outSignLength))))

	return C.GoString((*C.char)(cOutSign)), m.wrapError(rc)

}
