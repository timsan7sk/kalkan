package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

unsigned long get_time_from_sign(char *inData, int inDataLength, int flags, int inSigId, time_t *outDateTime) {
    return kc_funcs->KC_GetTimeFromSig(inData, inDataLength, flags, inSigId, outDateTime);
}
*/
import "C"
import (
	"time"
	"unsafe"
)

// Obtains the signature time.
func (m *Module) GetTimeFromSign(cms string, sigID int, flag Flag) (time.Time, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()
	// preparing variable and freeing memory when finished.
	cCMS := C.CString(cms)
	defer C.free(unsafe.Pointer(cCMS))

	outDateTime := C.time_t(0)
	rc := int(C.get_time_from_sign(cCMS, C.int(len(cms)), C.int(int(flag)), C.int(sigID), &outDateTime))

	// checking for errors.
	if err := m.wrapError(rc); err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(outDateTime), 0).UTC(), nil
}
