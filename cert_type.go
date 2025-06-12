package kalkan

// Constants defining the ownership of a certificate.
const (
	CertTypeCA           = 0x00000201     // Root CA certificate.
	CertTypeIntermediate = 0x00000202     // Intermediate CA certificate.
	CertTypeUser         = 0x00000204     // User certificate.
	CertTypeIndividual   = "individual"   // Certificate of an individual.
	CertTypeOrganization = "organization" // Certificate of legal entity.
)
