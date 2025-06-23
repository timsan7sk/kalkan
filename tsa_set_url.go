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
	"net/url"
	"unsafe"
)

// Sets the TSA service address.
func (m *Module) TSASetURL(tsaURL string) error {
	// checking URL is valid.
	_, err := url.ParseRequestURI(tsaURL)
	if err != nil {
		return err
	}
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()
	// preparing variable and freeing memory when finished.
	cURL := C.CString(tsaURL)
	defer C.free(unsafe.Pointer(cURL))
	// set up a TSA service address.
	C.set_tsa_url(cURL)
	return nil
}
