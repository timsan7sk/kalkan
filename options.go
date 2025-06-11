package kalkan

type Options struct {
	OCSP  string  `json:"ocsp"`  // URL of the online service for checking the certificate status.
	TSP   string  `json:"tsp"`   // URL of the timestamp service.
	Certs []Certs `json:"certs"` // Root CA certificates and intermediate CA certificates.
}

// Gets module options.
func (m *Module) Options() Options {
	return m.o
}
