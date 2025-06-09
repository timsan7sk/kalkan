package kalkan

// Determines the ownership of the certificate.
type CertType int

// Constants defining the ownership of a certificate.
const (
	CertTypeCA           CertType = 0x00000201 // Root CA certificate.
	CertTypeIntermediate CertType = 0x00000202 // Intermediate CA certificate.
	CertTypeUser         CertType = 0x00000204 // User certificate.
)
