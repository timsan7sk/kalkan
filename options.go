package kalkan

const (
	ProdOCSP = "http://ocsp.pki.gov.kz" // Production OCSP Service.
	ProdTSP  = "http://tsp.pki.gov.kz"  // Production TSP Service.

	TestOCSP = "http://test.pki.gov.kz/ocsp/" // Test OCSP Service.
	TestTSP  = "http://test.pki.gov.kz/tsp/"  // Test TSP Service.
)

type Options struct {
	OCSP      string  `json:"ocsp"`       // URL of the online service for checking the certificate status.
	TSP       string  `json:"tsp"`        // URL of the timestamp service.
	Certs     []Certs `json:"certs"`      // Root CA certificates and intermediate CA certificates.
	LoadCerts bool    `json:"load_certs"` // Do need to load root CA certificates and intermediate CA certificates.
}

// Obtains module options.
func (m *Module) Options() Options {
	return m.o
}

func (o *Options) setDefaults() {
	if o == nil {
		return
	}

	if o.TSP == "" {
		o.TSP = ProdTSP
	}

	if o.OCSP == "" {
		o.OCSP = ProdOCSP
	}

	o.LoadCerts = true
}
