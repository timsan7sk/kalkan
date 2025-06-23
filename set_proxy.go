package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long set_proxy(int flags, char *inProxyAddr, char *inProxyPort, char *inUser, char *inPass) {
	return kc_funcs->KC_SetProxy(flags, inProxyAddr, inProxyPort, inUser, inPass);
}
*/
import "C"
import (
	"fmt"
	"net/url"
	"unsafe"
)

// Set proxy server settings.
func (m *Module) SetProxy(flag Flag, url *url.URL) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if url == nil {
		return fmt.Errorf("%s", "URL must not be nil.")
	}
	cAddr := C.CString(url.Hostname())
	defer C.free(unsafe.Pointer(cAddr))

	cPort := C.CString(url.Port())
	defer C.free(unsafe.Pointer(cPort))

	cUser := C.CString(url.User.Username())
	defer C.free(unsafe.Pointer(cUser))

	pass, _ := url.User.Password()
	cPass := C.CString(pass)
	defer C.free(unsafe.Pointer(cPass))

	rc := int(C.set_proxy(C.int(int(flag)), cAddr, cPort, cUser, cPass))

	return m.wrapError(rc)

}
