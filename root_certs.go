package kalkan

import "crypto/x509"

type Certs struct {
	Cert *x509.Certificate
	Type CertType
}
