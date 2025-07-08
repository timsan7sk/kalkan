package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

int x509_certificate_get_info(char *inCert, int inCertLength, int propId, char *outData, int *outDataLength) {
    return kc_funcs->X509CertificateGetInfo(inCert, inCertLength, propId, (unsigned char*)outData, outDataLength);
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// Obtains field/extension values ​​from a certificate.
//
// The certificate must be previously loaded using one of the following functions:
//
// - LoadKeyStore()
//
// - X509LoadCertificateFromFile()
//
// - X509LoadCertificateFromBuffer().
func (m *Module) X509CertificateGetInfo(inCert string, prop CertProp) (string, error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()

	// preparing variables and freeing memory when finished.
	cInCert := C.CString(inCert)
	defer C.free(unsafe.Pointer(cInCert))
	outDataLength := 32768
	data := C.malloc(C.ulong(C.sizeof_char * outDataLength))
	defer C.free(data)

	// getting info from certificate.
	rc := int(C.x509_certificate_get_info(cInCert, C.int(len(inCert)), C.int(int(prop)), (*C.char)(data), (*C.int)(unsafe.Pointer(&outDataLength))))

	// return result and checking for errors.
	return C.GoString((*C.char)(data)), m.wrapError(rc)
}

func (m *Module) X509CertificateGetSummary(inCert string) (s Summary, err error) {
	// locking the module and unlocking it after completion.
	m.mu.Lock()
	defer m.mu.Unlock()
	v := reflect.ValueOf(s)
	p := reflect.ValueOf(&s)
	e := p.Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if k := e.Field(i).Kind(); k == reflect.String {
			fmt.Println("String ", t.Field(i).Name)

		}
		if field.Anonymous { // field.Anonymous indicates an embedded field
			// if it's an embedded struct, iterate its fields
			eVal := value
			eTyp := eVal.Type()

			for j := 0; j < eTyp.NumField(); j++ {
				eField := eTyp.Field(j)
				// eValue := eVal.Field(j)
				// _ = eValue.Elem()
				if l := len(eField.Tag.Get("cert_prop")); l > 0 {
					// base 0 infers base from prefix
					_, err := strconv.ParseInt(eField.Tag.Get("cert_prop"), 0, 64)
					if err != nil {
						return Summary{}, err
					}
					// embeddedFieldValue := embeddedVal.Field(j)
				}
			}
		}
		// fmt.Printf("Field: %+v, Tag: %+v\n", field.Name, field.Tag.Get("cert_prop"))

	}
	// t := p.Type()
	// for i := 0; i < t.NumField(); i++ {
	// 	// check if the field is an embedded struct
	// 	if field.Anonymous { // field.Anonymous indicates an embedded field
	// 		// if it's an embedded struct, iterate its fields
	// 		eVal := value
	// 		eTyp := eVal.Type()

	// 		for j := 0; j < eTyp.NumField(); j++ {
	// 			eField := eTyp.Field(j)
	// 			// eValue := eVal.Field(j)
	// 			// _ = eValue.Elem()
	// 			if l := len(eField.Tag.Get("cert_prop")); l > 0 {
	// 				// base 0 infers base from prefix
	// 				intTag, err := strconv.ParseInt(eField.Tag.Get("cert_prop"), 0, 64)
	// 				if err != nil {
	// 					return Summary{}, err
	// 				}
	// 				// embeddedFieldValue := embeddedVal.Field(j)
	// 				fmt.Printf("Embedded Field: %s, Tag: %08x\n", eField.Name, intTag)
	// 			}
	// 		}
	// 	} else {
	// 		fmt.Printf("Field: %s, Value: %v\n", field.Name, value.Interface())
	// 	}
	// }
	return Summary{}, nil
}
