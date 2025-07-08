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
	err := mod.X509LoadCertificateFromBuffer(testCertGOST2, kalkan.CertEncodingTypeB64)
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
func TestX509CertificateGetSummary(t *testing.T) {
	_, err := mod.X509CertificateGetSummary("")
	if err != nil {
		t.Fatal(err)
	}
}
func TestX509CertificateGetInfo(t *testing.T) {
	_, err := mod.X509CertificateGetInfo(testCertGOST2, kalkan.CertPropCertCN)
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidateCert(t *testing.T) {
	_, err := mod.X509ValidateCertificate(testCertGOST2, kalkan.ValidateTypeOCSP, mod.Options().OCSP)
	if err != nil {
		t.Log(err)
	}
}
