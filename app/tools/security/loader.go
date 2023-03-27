package security

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func GenerateTLSConfig(certBytes, keyBytes, caCertBytes []byte) (*tls.Config, error) {
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	if certPEM == nil {
		return nil, errors.New("failed to encode certificate")
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyBytes,
	})
	if keyPEM == nil {
		return nil, errors.New("failed to encode private key")
	}

	serverCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}

	caCert, err := x509.ParseCertificate(caCertBytes)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(caCert)

	return &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		RootCAs:      certPool.Clone(),
		ClientCAs:    certPool.Clone(),
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}, nil
}
