package ca

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

func GenerateCACertificate() (cert []byte, key []byte, err error) {
	caPrivateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	caTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "openPanel CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(25, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            2,
	}
	cert, err = x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return nil, nil, err
	}
	key, err = x509.MarshalPKCS8PrivateKey(caPrivateKey)
	if err != nil {
		return nil, nil, err
	}
	return cert, key, nil
}

func GenerateCertificateSigningRequest(commonName string) (csr []byte, key []byte, err error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	template := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{"openPanel"},
		},
		SignatureAlgorithm: x509.ECDSAWithSHA384,
	}
	csr, err = x509.CreateCertificateRequest(rand.Reader, &template, privateKey)
	if err != nil {
		return nil, nil, err
	}
	key, err = x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, nil, err
	}
	return csr, key, nil
}

func SignCsr(caCertBytes, caKeyBytes, csrBytes []byte) ([]byte, error) {
	caKey, err := x509.ParsePKCS8PrivateKey(caKeyBytes)
	if err != nil {
		return nil, err
	}
	caCert, err := x509.ParseCertificate(caCertBytes)
	if err != nil {
		return nil, err
	}
	csr, err := x509.ParseCertificateRequest(csrBytes)
	if err != nil {
		return nil, err
	}

	randSerial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 80))
	if err != nil {
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: randSerial,
		Subject:      csr.Subject,
		NotBefore:    time.Now(),
		NotAfter:     caCert.NotAfter,
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}
	cert, err := x509.CreateCertificate(rand.Reader, &template, caCert, csr.PublicKey, caKey)
	if err != nil {
		return nil, err
	}
	return cert, nil
}
