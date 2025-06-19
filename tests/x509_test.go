package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

func TestX509ExportCertificateFromStrore(t *testing.T) {
	_, err := mod.X509ExportCertificateFromStore("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestX509ExportCertificateFromBuffer(t *testing.T) {
	err := mod.X509LoadCertificateFromBuffer([]byte(testCertGOST2), kalkan.CertEncodingTypeB64)
	if err != nil {
		t.Fatal(err)
	}
}

func TestX509LoadCertificateFromFile(t *testing.T) {
	err := mod.X509LoadCertificateFromFile("GOST512.crt", kalkan.CertTypeUser)
	if err != nil {
		t.Fatal(err)
	}
}
