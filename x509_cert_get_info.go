package kalkan

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "KalkanCrypt.h"

static inline int x509_certificate_get_info(char *inCert, int inCertLength, int propId, char *outData, int *outDataLength) {
    if (kc_funcs == NULL || kc_funcs->X509CertificateGetInfo == NULL) {
        return -1;
    }
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

// X509CertificateGetInfo извлекает сырые данные из сертификата.
// Реализована защита от Buffer Overread и оптимизированное управление памятью через сборщик мусора Go.
func (m *Module) X509CertificateGetInfo(inCert string, prop CertProp) (string, error) {
	if !prop.IsExist() {
		return "", fmt.Errorf("invalid or unsupported property ID: %d", prop)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	cInCert := C.CString(inCert)
	defer C.free(unsafe.Pointer(cInCert))

	// Выделяем буфер в памяти Go. Это безопаснее и быстрее, чем C.malloc.
	// CGO разрешает передавать указатель на память Go в Си-функции, если они не сохраняют этот указатель асинхронно.
	buf := make([]byte, 32768)
	outLen := C.int(len(buf))

	// Вызов C-моста, передаем адрес базового массива слайса
	rc := int(C.x509_certificate_get_info(
		cInCert,
		C.int(len(inCert)),
		C.int(prop),
		(*C.char)(unsafe.Pointer(&buf[0])),
		&outLen,
	))

	// Делегируем проверку ошибок встроенному механизму (заменяет прямую проверку C.KCR_OK)
	if err := m.wrapError(rc); err != nil {
		return "", err
	}

	actualLen := int(outLen)
	if actualLen <= 0 {
		return "", nil
	}

	// Защита от Buffer Overread: используем C.GoStringN с точным фактическим размером,
	// игнорируя неинициализированный хвост буфера.
	return C.GoStringN((*C.char)(unsafe.Pointer(&buf[0])), C.int(actualLen)), nil
}

// X509CertificateGetSummary заполняет структуру Summary с использованием рефлексивного обхода (AST-like traversal).
func (m *Module) X509CertificateGetSummary(inCert string) (*Summary, error) {
	s := &Summary{}

	// Инициация рекурсивного обхода (Reflection Engine)
	m.populateTags(reflect.ValueOf(s).Elem(), inCert)

	s.Init()
	return s, nil
}

// populateTags — рекурсивная функция обхода структуры. Поддерживает композицию (Anonymous/Embedded Fields).
func (m *Module) populateTags(v reflect.Value, inCert string) {
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		// Если поле является встраиваемой структурой (например, CertSubject), спускаемся рекурсивно.
		if fieldType.Anonymous && fieldValue.Kind() == reflect.Struct {
			m.populateTags(fieldValue, inCert)
			continue
		}

		tag := fieldType.Tag.Get("cert_prop")
		if tag == "" {
			continue
		}

		n, err := strconv.ParseInt(tag, 0, 64)
		if err != nil {
			continue
		}

		cp := CertProp(n)
		if cp.IsExist() {
			val, err := m.X509CertificateGetInfo(inCert, cp)

			// Атомарная установка значения с отсечением префиксов (TrimAtEqual).
			// Использование TrimAtEqual на уровне бизнес-логики сохраняет чистоту низкоуровневой X509CertificateGetInfo.
			if err == nil && val != "" && fieldValue.CanSet() && fieldValue.Kind() == reflect.String {
				fieldValue.SetString(TrimAtEqual(val))
			}
		}
	}
}
