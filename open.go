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
	m := Module{
		h:  h,
		mu: sync.Mutex{},
	}
	m.o.setDefaults()
	return &m, nil
}
