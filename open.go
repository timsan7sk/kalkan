package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdio.h>
#include "KalkanCrypt.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// Library name.
// var libName = getLibraryName()

// Module struct for interaction with the KalkanCrypt library.
type Module struct {
	h  unsafe.Pointer // handle to a library (.so).
	mu sync.Mutex     // mutual exclusion lock.
}

// Obtains the new client.
func Open(libName string) (*Module, error) {
	h, err := dlOpen(libName)
	if err != nil {
		return nil, err
	}
	return &Module{
		h:  h,
		mu: sync.Mutex{},
	}, nil
}
