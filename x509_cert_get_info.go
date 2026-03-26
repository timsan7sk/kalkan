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
	"log"
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
	st := reflect.TypeOf(s)
	// in := regexp.MustCompile(`[0-9]{12}$`)
	// sn := regexp.MustCompile(`[A-Z0-9]{40}$`)
	t := v.Type()
	for i := 0; i < st.NumField(); i++ {
		if !st.Field(i).Anonymous {
			if c := t.Field(i).Tag.Get("cert_prop"); c != "" {
				n, err := strconv.ParseInt(t.Field(i).Tag.Get("cert_prop"), 0, 64)
				if err != nil {
					log.Println(err)
				}
				if cp := CertProp(n); cp.IsExist() {
					ttt, err := m.X509CertificateGetInfo(inCert, cp)
					if err != nil {
						log.Println(err)
					} else {
						w := TrimAtEqual(ttt)
						p.Elem().Field(i).SetString(w)
					}
				}
			}

		}
	}
	log.Printf("FieldName: %+v", p)
	// for i := 0; i < t.NumField(); i++ {
	// field := t.Field(i)

	// value := v.Field(i)
	// pf := p.Elem().Field(i)
	// if k := e.Field(i).Kind(); k == reflect.String {
	// for i := 0; i < reflect.ValueOf(s).NumField(); i++ {
	// 	log.Printf("%s %s\n", reflect.ValueOf(s).Type().Field(i).Name, reflect.ValueOf(s).Field(i).Interface())

	// }
	// }
	// if field.Anonymous {
	// 	eTyp := value.Type()
	// 	for j := 0; j < eTyp.NumField(); j++ {
	// 		eField := eTyp.Field(j)
	// 		if len(eField.Tag.Get("cert_prop")) > 0 {
	// 			// log.Println("PeField.Tag.Get:", eField.Tag.Get("cert_prop"))

	// 			n, err := strconv.ParseInt(eField.Tag.Get("cert_prop"), 0, 64)
	// 			// log.Printf("FieldName: %s, Tag: %d\n", eField.Name, n)
	// 			if err != nil {
	// 				log.Println("ParseInt:", err)
	// 			}
	// 			if cp := CertProp(n); cp.IsExist() {
	// 				vvv, err := m.X509CertificateGetInfo(inCert, cp)
	// 				if err != nil {
	// 					continue
	// 					// log.Println("X509CertificateGetInfo: ", err)
	// 				} else {
	// 					pf.Field(j).SetString(vvv)
	// 				}
	// 				// number := in.FindString(vvv)
	// 			}
	// 		}
	// 	}
	// }
	// }

	// for i := 0; i < reflect.ValueOf(s).NumField(); i++ {
	// 	log.Printf("Field: %s\tValue: %v\n", reflect.ValueOf(s).Type().Field(i).Name, reflect.ValueOf(s).Field(i).Interface())

	// }
	return Summary{}, nil
}
