package kalkan

/*
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <dlfcn.h>
*/
import "C"
import "fmt"

// Closes a Module.
func (m *Module) Close() error {
	C.dlerror()
	C.dlclose(m.h)
	if err := C.dlerror(); err != nil {
		return fmt.Errorf("%s", C.GoString(err))
	}
	return nil
}
