package ca

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func generateTestCertificate() ([]byte, []byte, []byte, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, nil, err
	}

	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2022),
		Subject: pkix.Name{
			CommonName: "Test CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24),
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &privKey.PublicKey, privKey)
	if err != nil {
		return nil, nil, nil, err
	}

	serverCert := &x509.Certificate{
		SerialNumber: big.NewInt(123456789),
		Subject: pkix.Name{
			CommonName: "localhost",
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(time.Hour * 24),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	serverCertBytes, err := x509.CreateCertificate(rand.Reader, serverCert, ca, &privKey.PublicKey, privKey)
	if err != nil {
		return nil, nil, nil, err
	}

	keyBytes := x509.MarshalPKCS1PrivateKey(privKey)

	return caBytes, serverCertBytes, keyBytes, nil
}

func TestGenerateRPCTLSConfig(t *testing.T) {
	caCertBytes, certBytes, keyBytes, err := generateTestCertificate()
	assert.NoError(t, err, "generateTestCertificate failed")

	tlsConfig, err := GenerateRPCTLSConfig(certBytes, keyBytes, caCertBytes)
	assert.NoError(t, err, "GenerateRPCTLSConfig failed")
	assert.NotNil(t, tlsConfig, "TLS Config should not be nil")
	assert.Equal(t, tls.RequireAndVerifyClientCert, tlsConfig.ClientAuth, "Wrong ClientAuth value")
}

func TestGenerateHTTPTLSConfig(t *testing.T) {
	_, certBytes, keyBytes, err := generateTestCertificate()
	assert.NoError(t, err, "generateTestCertificate failed")

	tlsConfig, err := GenerateHTTPTLSConfig(certBytes, keyBytes)
	assert.NoError(t, err, "GenerateHTTPTLSConfig failed")
	assert.NotNil(t, tlsConfig, "TLS Config should not be nil")
	assert.Equal(t, tls.NoClientCert, tlsConfig.ClientAuth, "Wrong ClientAuth value")
}