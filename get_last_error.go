package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long getLastError() {
//     return kc_funcs->KC_GetLastError();
// }
import "C"

// Obtains last error code.
func (m *Module) GetLastError() ErrorCode {
	return ErrorCode(C.getLastError())
}
