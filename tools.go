package kalkan

import (
	"encoding/xml"
	"strings"
)

// Trims the slice.
func trimSlice(data []byte) []byte {
	for i, v := range data {
		if v == 0 {
			return data[:i]
		}
	}
	return data
}

// Wraps the XML document in SOAP format, or more precisely writes the contents under the soap:Body tag.
func WrapSOAP(inData, id string) (string, error) {
	// Creating SOAP wrapper.
	e := soapEnvelope{
		SOAP: xmlnsSOAP,
		WSU:  xmlnsWSU,
		Body: soapBody{
			ID:      id,
			Content: replace,
		},
	}
	// Marhaling SOAP wrapper.
	b, err := xml.Marshal(e)
	if err != nil {
		return "", err
	}
	// Adding data to the soap:Body tag and return result.
	return strings.Replace(string(b), replace, inData, 1), nil
}
