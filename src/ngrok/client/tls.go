package client

import (
	_ "crypto/sha512"
	"crypto/tls"
	"crypto/x509"
	//"encoding/pem"
	//"fmt"
	"io/ioutil"
	//"ngrok/client/assets"
)

/*
func LoadTLSConfig(rootCertPaths []string) (*tls.Config, error) {
	pool := x509.NewCertPool()

	for _, certPath := range rootCertPaths {
		rootCrt, err := assets.Asset(certPath)
		if err != nil {
			return nil, err
		}

		pemBlock, _ := pem.Decode(rootCrt)
		if pemBlock == nil {
			return nil, fmt.Errorf("Bad PEM data")
		}

		certs, err := x509.ParseCertificates(pemBlock.Bytes)
		if err != nil {
			return nil, err
		}

		pool.AddCert(certs[0])
	}

	return &tls.Config{RootCAs: pool}, nil
}
*/

func LoadTLSConfig(crtPath string) (*tls.Config, error) {
	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(crtPath)
	if err != nil {
		return nil, err
	}
	pool.AppendCertsFromPEM(caCrt)

	return &tls.Config{RootCAs: pool}, nil
}

