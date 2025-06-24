package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long get_token(unsigned long storage, char *tokens, unsigned long *tk_count) {
	return kc_funcs->KC_GetTokens(storage, tokens, tk_count);
}
*/
import "C"
import "unsafe"

// Obtains a string of connected storage devices.
func (m *Module) GetToken(store StoreType) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	cToken := C.malloc(C.ulong(C.sizeof_char * 8192))
	defer C.free(cToken)
	count := uint64(0)

	rc := int(C.get_token(C.ulong(uint(store)), (*C.char)(cToken), (*C.ulong)(unsafe.Pointer(&count))))

	err := m.wrapError(rc)
	if err != nil {
		return "", err
	}

	return C.GoString((*C.char)(cToken)), nil

}
