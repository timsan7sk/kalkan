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
)

// Obtains the pointer to the Module.
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
