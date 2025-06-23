package kalkan

import (
	"net/url"
	"time"
)

// Wrapper over KalkanCrypt library methods.
type Kalkan interface {
	Init() error

	LoadKeyStore(path, pwd, alias string, storeType StoreType) error

	HashData(alg HashAlg, flags Flag, inData string) (string, error)

	SignData(inSign, inData, alias string, flag Flag) (string, error)
	SignXML(xml, alias string, flags Flag, signNodeID, parentSignNode, parentNameSpace string) (string, error)
	SignWSSE(alias, inData, signNodeID string, flags Flag) (string, error)

	VerifyData(inSign, inData, alias string, flags Flag) (*VerifiedData, error)
	VerifyXML(alias, inData string, flags Flag) (string, error)

	GetCertFromCMS(cms string, signID int, flag Flag) (string, error)
	GetCertFromXML(xml string, signID int) (string, error)
	GetCertList() (string, int, error)

	GetTimeFromSign(cms string, sigID int, flag Flag) (time.Time, error)

	GetTokens(store StoreType) (string, int, error)

	GetLastError() ErrorCode
	GetLastErrorString() (ErrorCode, string)

	GetSignAlgorithmFromXML(xml string) (string, error)

	SetProxy(flag Flag, url *url.URL) error
	TSASetURL(tsaURL string) error

	X509CertificateGetInfo(inCert string, prop CertProp) (string, error)
	X509ExportCertificateFromStore(alias string) (string, error)
	X509LoadCertificateFromBuffer(inCert []byte, flags CertEncodingType) error
	X509LoadCertificateFromFile(certPath string, certType CertType) error
	X509ValidateCertificate(inCert string, validateType ValidateType, validatePath string) (string, error)

	XMLFinalize()
	Finalize()
	Close() error
}

var _ Kalkan = (*Module)(nil)
