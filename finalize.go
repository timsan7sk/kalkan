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

func (m *Module) Finalize() {
	C.finalize()
}
