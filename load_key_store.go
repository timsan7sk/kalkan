package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

	load_key_store(int storage, char *password, int passLen, char *container, int containerLen, char *alias) {
		return kc_funcs->KC_LoadKeyStore(storage, password, passLen, container, containerLen, alias);
	}
*/
import "C"
import "unsafe"

// Loads keys/certificates from storage.
func (m *Module) LoadKeyStore(path, pwd, alias string, storeType StoreType) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	cPwd := C.CString(pwd)
	defer C.free(unsafe.Pointer(cPwd))

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	rc := int(C.load_key_store(C.int(int(storeType)), cPwd, C.int(len(pwd)), cPath, C.int(len(path)), cAlias))

	return m.wrapError(rc)
}
