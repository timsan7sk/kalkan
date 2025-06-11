package kalkan

import (
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
)

type Certs struct {
	Cert *x509.Certificate
	Type CertType
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
