package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

void xml_finalize() {
	kc_funcs->KC_XMLFinalize();
}
*/
import "C"

// Frees memory and terminates the library with modules responsible for parsing,
// signing and checking data in XML format. No need to call each time when signing.
// Can only be called once after the cycle of signing XML files.
func (m *Module) XMLFinalize() {
	C.xml_finalize()
}
