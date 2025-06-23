package kalkan

/*
#include <dlfcn.h>
#include <stdio.h>
#include "KalkanCrypt.h"

	void get_function_list(void *f) {
	    void (*KC_GetFunctionList)(stKCFunctionsType **);
	    KC_GetFunctionList = (void (*)(stKCFunctionsType **))f;
	    KC_GetFunctionList(&kc_funcs);
	}

	int init() {
	    return (kc_funcs)->KC_Init();
	}
*/
import "C"

// Initializes the library.
func (m *Module) Init() error {
	f, err := dlSym(m.h)
	if err != nil {
		return err
	}
	C.get_function_list(f)
	rv := int(C.init())
	return m.wrapError(rv)
}
