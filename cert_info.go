package kalkan

import (
	"strings"
)

const (
	OwnerIndividual   = "individual"   // Certificate of an individual.
	OwnerOrganization = "organization" // Certificate of legal entity.
)

// see: https://adilet.zan.kz/rus/docs/V2000021440
const (
	oidSubjectIndividual      = "1.2.398.3.3.4.1.1"
	oidSubjectRoleCEO         = "1.2.398.3.3.4.1.2.1"
	oidSubjectRoleSign        = "1.2.398.3.3.4.1.2.2"
	oidSubjectRoleSignFinance = "1.2.398.3.3.4.1.2.3"
	oidSubjectRoleHR          = "1.2.398.3.3.4.1.2.4"
	oidSubjectRoleEmployee    = "1.2.398.3.3.4.1.2.5"
)

const (
	CertSubjectRoleUndefined CertSubjectRole = iota
	CertSubjectRoleCEO
	CertSubjectRoleSign
	CertSubjectRoleSignFinance
	CertSubjectRoleHR
	CertSubjectRoleEmployee
)

type (
	Summary struct {
		Owner string `cert_prop:"0x00000815"` // Defines the form of ownership of the certificate owner.
		CertSubject
		CertOrganization
		CertIssuer
		PublicKey    string `cert_prop:"0x0000081d"` // CertPropPubKey
		SerialNumber string `cert_prop:"0x00000819"` // CertPropCertSN
		OCSP         string `cert_prop:"0x0000081f"` // CertPropOCSP
	}

	CertSubjectRole int

	CertSubject struct {
		CommonName string `cert_prop:"0x0000080a"` // CertPropSubjectCommonName
		LastName   string `cert_prop:"0x0000080b"` // CertPropSubjectGivenName
		Country    string `cert_prop:"0x00000807"` // CertPropSubjectCountryName
		IIN        string `cert_prop:"0x0000080d"` // CertPropSubjectSerialNumber
		DN         string `cert_prop:"0x0000081b"` // CertPropSubjectDN
		Locality   string `cert_prop:"0x00000809"` // CertPropSubjectLocalityName
		OrgName    string `cert_prop:"0x0000080f"` // CertPropSubjectOrgName
		OrgUnit    string `cert_prop:"0x00000810"` // CertPropSubjectOrgUnitName
		BC         string `cert_prop:"0x00000811"` // CertPropSubjectBC
		DC         string `cert_prop:"0x00000812"` // CertPropSubjectDC
		Email      string `cert_prop:"0x0000080e"` // CertPropSubjectEmail
	}

	CertOrganization struct {
		OrgName     string `cert_prop:"0x0000080f"` // CertPropSubjectOrgName
		BIN         string `cert_prop:"0x00000810"` // CertPropSubjectOrgUnitName
		SubjectRole string `cert_prop:"0x0000081e"` // CertPropPoliciesID
	}

	CertIssuer struct {
		CommonName string `cert_prop:"0x00000806"` // CertPropIssuerCommonName
		Country    string `cert_prop:"0x00000801"` // CertPropIssuerCountryName
		Locality   string `cert_prop:"0x00000803"` // CertPropIssuerLocalityName
		OrgName    string `cert_prop:"0x00000805"` // CertPropIssuerOrgUnitName
		OrgUnit    string `cert_prop:"0x00000804"` // CertPropIssuerOrgName
		DN         string `cert_prop:"0x0000081a"` // CertPropIssuerDN
	}
)

func (s *Summary) Init() bool {
	if s.BIN != "" {
		s.Owner = OwnerOrganization
	} else {
		s.Owner = OwnerIndividual
	}
	return true
}

// TrimAtEqual извлекает значение после знака '='.
// Алгебраическая оптимизация: использование IndexByte исключает аллокации строк (в отличие от Split).
func TrimAtEqual(s string) string {
	if idx := strings.IndexByte(s, '='); idx >= 0 {
		return s[idx+1:]
	}
	return s
}
