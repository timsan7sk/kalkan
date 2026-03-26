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
	_, err := mod.X509CertificateGetSummary(testCertGOST2)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(c)
}
func TestX509CertificateGetInfo(t *testing.T) {
	c := testGetCertFromXML(t)
	// t.Logf("Arg: %s\n", c[:516])
	a, err := mod.X509CertificateGetInfo(c, kalkan.CertPropNotAfter)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Arg: %s\n", a)
}

func TestValidateCert(t *testing.T) {
	_, err := mod.X509ValidateCertificate(testCertGOST2, kalkan.ValidateTypeOCSP, mod.Options().OCSP)
	if err != nil {
		t.Log(err)
	}
}
