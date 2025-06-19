package kalkan

const (
	prodOCSP = "http://ocsp.pki.gov.kz"       // Production OCSP Service.
	prodTSP  = "http://tsp.pki.gov.kz"        // Production TSP Service.
	testOCSP = "http://test.pki.gov.kz/ocsp/" // Test OCSP Service.
	testTSP  = "http://test.pki.gov.kz/tsp/"  // Test TSP Service.
)

var (
	ProdOpts = []Option{
		setOCSP(prodOCSP),
		setTSP(prodTSP),
	}
	TestOpts = []Option{
		setOCSP(testOCSP),
		setTSP(testTSP),
	}
)

type Option func(o *Options)

type Options struct {
	OCSP      string  `json:"ocsp"`       // URL of the online service for checking the certificate status.
	TSP       string  `json:"tsp"`        // URL of the timestamp service.
	Certs     []Certs `json:"certs"`      // Root CA certificates and intermediate CA certificates.
	LoadCerts bool    `json:"load_certs"` // Do need to load root CA certificates and intermediate CA certificates.
}

// Gets module options.
func (m *Module) Options() Options {
	return m.o
}

// Sets default options.
func (o *Options) setDefaults() {
	if o == nil {
		return
	}
	if o.TSP == "" {
		o.TSP = prodTSP
	}
	if o.OCSP == "" {
		o.OCSP = prodOCSP
	}
	o.LoadCerts = true
}

// Sets TSP URL.
func setTSP(s string) Option {
	return func(o *Options) {
		o.TSP = s
	}
}

// Sets OCSP URL.
func setOCSP(s string) Option {
	return func(o *Options) {
		o.OCSP = s
	}
}
