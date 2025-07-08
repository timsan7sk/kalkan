package kalkan

import (
	"time"
)

const (
	OwnerIndividual   Owner = "individual"   // Certificate of an individual.
	OwnerOrganization Owner = "organization" // Certificate of legal entity.
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
		Owner Owner
		CertSubject
		CertOrganization
		CertIssuer
		PublicKey    string    `cert_prop:"0x0000081d"` // CertPropPubKey
		SerialNumber string    `cert_prop:"0x00000819"` // CertPropCertCN
		NotAfter     time.Time `cert_prop:"0x00000814"` // CertPropNotAfter
		NotBefore    time.Time `cert_prop:"0x00000813"` // CertPropNotBefore
	}
	Owner           string // Defines the form of ownership of the certificate owner.
	CertSubjectRole int
	CertSubject     struct {
		CommonName string `cert_prop:"0x0000080a"` // CertPropSubjectCommonName
		LastName   string `cert_prop:"0x0000080b"` // CertPropSubjectGivenName
		Country    string `cert_prop:"0x00000807"` // CertPropSubjectCountryName
		IIN        string `cert_prop:"0x0000080d"` // CertPropSubjectSerialNumber
		DN         string `cert_prop:"0x0000081b"` // CertPropSubjectDN
	}
	CertOrganization struct {
		Name        string          `cert_prop:"0x0000080f"` // CertPropSubjectOrgName
		BIN         string          `cert_prop:"0x00000810"` // CertPropSubjectOrgUnitName
		SubjectRole CertSubjectRole `cert_prop:""`
	}
	CertIssuer struct {
		CommonName string `cert_prop:"0x00000806"` // CertPropIssuerCommonName
		Country    string `cert_prop:"0x00000801"` // CertPropIssuerCountryName
		DN         string `cert_prop:"0x0000081a"` // CertPropIssuerDN
	}
)

func (s *Summary) Init() bool {

	return true
}
