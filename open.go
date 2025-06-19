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
func Open(libName string, options ...Option) (*Module, error) {
	h, err := dlOpen(libName)
	if err != nil {
		return nil, err
	}
	m := Module{
		h:  h,
		mu: sync.Mutex{},
	}
	m.o.setDefaults()
	for _, op := range options {
		op(&m.o)
	}
	return &m, nil
}
