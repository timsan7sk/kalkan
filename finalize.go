package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

void finalize() {
	kc_funcs->KC_Finalize();
}
*/
import "C"

// Frees resources of the KalkanCryptCOM crypto provider and terminates the library.
func (m *Module) Finalize() {
	C.finalize()
}
