package kalkan

import (
	"sync"
	"unsafe"
)

// Module struct for interaction with the KalkanCrypt library.
type Module struct {
	h  unsafe.Pointer // handle to a library (.so).
	o  Options        // Options of the Module.
	mu sync.Mutex     // mutual exclusion lock.
}
