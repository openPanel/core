package ca

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/openPanel/core/app/constant"
)

func loadServerCert(certBytes, keyBytes []byte) (*tls.Certificate, error) {
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
	return &serverCert, nil
}

func GenerateRPCTLSConfig(certBytes, keyBytes, caCertBytes []byte) (*tls.Config, error) {
	serverCert, err := loadServerCert(certBytes, keyBytes)
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
		Certificates: []tls.Certificate{*serverCert},
		RootCAs:      certPool.Clone(),
		ClientCAs:    certPool.Clone(),
		ClientAuth:   tls.RequireAndVerifyClientCert,
		NextProtos:   []string{constant.RPCProtocolName},
		MinVersion:   tls.VersionTLS13,
	}, nil
}

func GenerateHTTPTLSConfig(certBytes, keyBytes []byte) (*tls.Config, error) {
	serverCert, err := loadServerCert(certBytes, keyBytes)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{*serverCert},
		ClientAuth:   tls.NoClientCert,
	}, nil
}
