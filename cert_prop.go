package kalkan

// Defines the value of a field/extension in a request/certificate.
type CertProp int

const ( // Constants defining the value of a field/extension in a request/certificate.
	CertPropIssuerCountryName   CertProp = 0x00000801 // Country of issuer.
	CertPropIssuerSOPN          CertProp = 0x00000802 // Name of the state or province of the issuer.
	CertPropIssuerLocalityName  CertProp = 0x00000803 // The issuer's locality name.
	CertPropIssuerOrgName       CertProp = 0x00000804 // Name of the issuer organization.
	CertPropIssuerOrgUnitName   CertProp = 0x00000805 // Name of the organizational unit of the issuer.
	CertPropIssuerCommonName    CertProp = 0x00000806 // Name and Surname of the issuer.
	CertPropSubjectCountryName  CertProp = 0x00000807 // Country of the subject.
	CertPropSubjectSOPN         CertProp = 0x00000808 // The name of the state or province of the subject.
	CertPropSubjectLocalityName CertProp = 0x00000809 // The subject's locality name.
	CertPropSubjectCommonName   CertProp = 0x0000080a // Common name of the subject.
	CertPropSubjectGivenName    CertProp = 0x0000080b // Subject given name.
	CertPropSubjectSurname      CertProp = 0x0000080c // Subject's surname.
	CertPropSubjectSerialNumber CertProp = 0x0000080d // Subject serial number.
	CertPropSubjectEmail        CertProp = 0x0000080e // Subject's email.
	CertPropSubjectOrgName      CertProp = 0x0000080f // Name of the entity organization.
	CertPropSubjectOrgUnitName  CertProp = 0x00000810 // Name of the organizational unit of the entity.
	CertPropSubjectBc           CertProp = 0x00000811 // Business category of the subject.
	CertPropSubjectDc           CertProp = 0x00000812 // Domain component of the subject.
	CertPropNotBefore           CertProp = 0x00000813 // Date valid from.
	CertPropNotAfter            CertProp = 0x00000814 // Date valid until.
	CertPropKeyUsage            CertProp = 0x00000815 // Using the key.
	CertPropExtKeyUsage         CertProp = 0x00000816 // Extended key usage.
	CertPropAuthKeyID           CertProp = 0x00000817 // Certification Authority Key Identifier.
	CertPropSubjKeyID           CertProp = 0x00000818 // Subject key identifier.
	CertPropCertCN              CertProp = 0x00000819 // Certificate serial number.
	CertPropIssuerDN            CertProp = 0x0000081a // Distinguishing name of the publisher.
	CertPropSubjectDN           CertProp = 0x0000081b // Distinguishing name of the subject.
	CertPropSignatureAlg        CertProp = 0x0000081c // Signature algorithm.
	CertPropPubKey              CertProp = 0x0000081d // Obtaining a public key.
	CertPropPoliciesID          CertProp = 0x0000081e // Obtaining the certificate policy identifier.
	CertPropOCSP                CertProp = 0x0000081f // Getting the OCSP URL.
	CertPropGetCRL              CertProp = 0x00000820 // Getting the CRL URL.
	CertPropGetDeltaCRL         CertProp = 0x00000821 // Getting the delta CRL URL.
)

//nolint:gochecknoglobals
var AllProps = []CertProp{
	CertPropIssuerCountryName,
	CertPropIssuerSOPN,
	CertPropIssuerLocalityName,
	CertPropIssuerOrgName,
	CertPropIssuerOrgUnitName,
	CertPropIssuerCommonName,
	CertPropSubjectCountryName,
	CertPropSubjectSOPN,
	CertPropSubjectLocalityName,
	CertPropSubjectCommonName,
	CertPropSubjectGivenName,
	CertPropSubjectSurname,
	CertPropSubjectSerialNumber,
	CertPropSubjectEmail,
	CertPropSubjectOrgName,
	CertPropSubjectOrgUnitName,
	CertPropSubjectBc,
	CertPropSubjectDc,
	CertPropNotBefore,
	CertPropNotAfter,
	CertPropKeyUsage,
	CertPropExtKeyUsage,
	CertPropAuthKeyID,
	CertPropSubjKeyID,
	CertPropCertCN,
	CertPropIssuerDN,
	CertPropSubjectDN,
	CertPropSignatureAlg,
	CertPropPubKey,
	CertPropPoliciesID,
	CertPropOCSP,
	CertPropGetCRL,
	// KCCertPropGetDeltaCRL,
}
