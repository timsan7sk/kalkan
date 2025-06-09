package kalkan

// Specifies the certificate encoding type.
type CertEncodingType int

// Constatns, specifies the certificate encoding type.
const (
	CertEncodingTypeDER CertEncodingType = 0x00000101 // Encoding DER
	CertEncodingTypePEM CertEncodingType = 0x00000102 // Encoding PEM
	CertEncodingTypeB64 CertEncodingType = 0x00000104 // Encoding Base64
)
