package kalkan

import (
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
)

type Certs struct {
	Cert *x509.Certificate
	Type int
}

// Data for downloading certificate.
type dataForLoadCert struct {
	url      string
	certType int
}

// Data for downloading certificates.
type dataForLoadCerts []dataForLoadCert

var (
	// Data for downloading production certificates.
	DataForProdCerts = dataForLoadCerts{
		{url: "https://pki.gov.kz/cert/root_rsa_2020.cer", certType: CertTypeCA},
		{url: "https://pki.gov.kz/cert/root_gost_2022.cer", certType: CertTypeCA},
		{url: "https://pki.gov.kz/cert/nca_rsa_2022.cer", certType: CertTypeIntermediate},
		{url: "https://pki.gov.kz/cert/nca_gost_2022.cer", certType: CertTypeIntermediate},
	}
	// Data for downloading test certificates.
	DataForTestCerts = dataForLoadCerts{
		{url: "https://test.pki.gov.kz/cert/root_gost_test.cer", certType: CertTypeCA},
		{url: "https://test.pki.gov.kz/cert/root_rsa_test.cer", certType: CertTypeCA},
		{url: "https://test.pki.gov.kz/cert/root_test_gost_2022.cer", certType: CertTypeCA},
		{url: "https://test.pki.gov.kz/cert/nca_gost_test.cer", certType: CertTypeIntermediate},
		{url: "https://test.pki.gov.kz/cert/nca_rsa_test.cer", certType: CertTypeIntermediate},
		{url: "https://test.pki.gov.kz/cert/nca_gost2022_test.cer", certType: CertTypeIntermediate},
	}
)

// Loads certificates from a URL and adds them to the module options.
func (m *Module) LoadRemoteCerts(data dataForLoadCerts) error {
	for i := range data {
		bytes, err := download(data[i].url)
		if err != nil {
			return err
		}
		cert, err := x509.ParseCertificate(bytes)
		if err != nil {
			return err
		}
		m.o.Certs = append(m.o.Certs, Certs{Cert: cert, Type: data[i].certType})
	}
	return nil
}

// Downloads the certificate from the URL to the store.
func download(url string) ([]byte, error) {
	// creating request.
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}
	// sends an HTTP request and returns an HTTP response.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// closing the body of the response.
	defer resp.Body.Close()
	// checking the response status.
	if resp.StatusCode != http.StatusOK {
		io.Copy(io.Discard, resp.Body)
		return nil, fmt.Errorf("status: %s, code: %d", resp.Status, resp.StatusCode)
	}
	// reading the body of the response.
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// return values.
	return bytes, nil
}
