package tests

import "testing"

func TestX509ExportCertificateFromStrore(t *testing.T) {
	gCert, err := mod.X509ExportCertificateFromStore("")
	if err != nil {
		t.Fatal(err)
	}
	if gCert != testCertGOST1 {
		t.Fatal("Cert mismatch")
	}
}
