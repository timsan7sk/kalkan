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

// Draft.
// TODO:
func (m *Module) X509CertificateGetSummary(inCert string) (s Summary, err error) {
	v := reflect.ValueOf(s)
	p := reflect.ValueOf(&s)
	// in := regexp.MustCompile(`[0-9]{12}$`)
	// sn := regexp.MustCompile(`[A-Z0-9]{40}$`)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		pf := p.Elem().Field(i)
		// if k := e.Field(i).Kind(); k == reflect.String {
		n, err := strconv.ParseInt(t.Field(i).Tag.Get("cert_prop"), 0, 64)
		if err != nil {
		}
		ttt, err := m.X509CertificateGetInfo(inCert, CertProp(n))
		if err != nil {
			fmt.Println(err)
		} else {
			p.Elem().Field(i).SetString(ttt)

		}
		// }
		if field.Anonymous {
			eTyp := value.Type()

			for j := 0; j < eTyp.NumField(); j++ {
				eField := eTyp.Field(j)
				if l := len(eField.Tag.Get("cert_prop")); l > 0 {

					n, err := strconv.ParseInt(eField.Tag.Get("cert_prop"), 0, 64)
					fmt.Printf("FielsName: %s, Tag: %d\n", eField.Name, n)
					if err != nil {
						fmt.Println(err)
					}
					vvv, err := m.X509CertificateGetInfo(inCert, CertProp(n))
					if err != nil {
						fmt.Println(err)
					}
					// number := in.FindString(vvv)
					pf.Field(j).SetString(vvv)
				}
			}
		}
	}
	fmt.Printf("%+v\n", s)
	return Summary{}, nil
}
