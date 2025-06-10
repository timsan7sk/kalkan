package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long get_tokens(unsigned long storage, char *tokens, unsigned long *tk_count) {
	return kc_funcs->KC_GetTokens(storage, tokens, tk_count);
}
*/
import "C"
import "unsafe"

// Obtains a string of connected storage devices and their number.
func (m *Module) GetTokens(store StoreType) (string, int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	cTokens := C.malloc(C.ulong(C.sizeof_char * 8192))
	defer C.free(cTokens)
	count := uint64(0)

	rc := int(C.get_tokens(C.ulong(uint(store)), (*C.char)(cTokens), (*C.ulong)(unsafe.Pointer(&count))))

	err := m.wrapError(rc)
	if err != nil {
		return "", 0, err
	}

	return C.GoString((*C.char)(cTokens)), int(count), nil

}
