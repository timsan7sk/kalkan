package kalkan

import "time"

const (
	OwnerIndividual   Owner = "individual"   // Certificate of an individual.
	OwnerOrganization Owner = "organization" // Certificate of legal entity.
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
		Owner        Owner
		Subject      CertSubject
		Organization *CertOrganization
		Issuer       CertIssuer
		PublicKey    string
		SerialNumber string
		NotAfter     time.Time
		NotBefore    time.Time
	}
	Owner           string // Defines the form of ownership of the certificate owner.
	CertSubjectRole int
	CertSubject     struct {
		CommonName string
		LastName   string
		Country    string
		IIN        string
		DN         string
	}
	CertOrganization struct {
		Name        string
		BIN         string
		SubjectRole CertSubjectRole
	}
	CertIssuer struct {
		CommonName string
		Country    string
		DN         string
	}
)
