package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

void set_tsa_url(char *tsaurl) {
    kc_funcs->KC_TSASetUrl(tsaurl);
}
*/
import "C"
import (
	"unsafe"
)

// Sets the TSA service address.
func (m *Module) TSASetURL(url string) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()
	// preparing variable and freeing memory when finished.
	cTSA := C.CString(url)
	defer C.free(unsafe.Pointer(cTSA))
	// set up a TSA service address.
	C.set_tsa_url(cTSA)
}
