package kalkan

// Certificate validation type.
type ValidateType int

// Constants defining the validation type.
const (
	ValidateTypeNothing ValidateType = 0x00000401 // Do nothing.
	ValidateTypeCRL     ValidateType = 0x00000402 // Check against the list of revoked certificates.
	ValidateTypeOCSP    ValidateType = 0x00000404 // Check the certificate using the OCSP service.
)
