package security

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"
)

func TestCertificateFunctions(t *testing.T) {
	// Test GenerateCACertificate
	caCert, caKey, err := GenerateCACertificate()
	if err != nil {
		t.Fatalf("Failed to generate CA certificate: %v", err)
	}

	// Test GenerateCertificateSigningRequest
	commonName := "example.com"
	csr, key, err := GenerateCertificateSigningRequest(commonName)
	if err != nil {
		t.Fatalf("Failed to generate CSR: %v", err)
	}

	// Test SignCsr
	signedCert, err := SignCsr(caCert, caKey, csr)
	if err != nil {
		t.Fatalf("Failed to sign CSR: %v", err)
	}

	// Check if the signed certificate has the correct fields
	signedCertParsed, err := x509.ParseCertificate(signedCert)
	if err != nil {
		t.Fatalf("Failed to parse signed certificate: %v", err)
	}

	if signedCertParsed.Subject.CommonName != commonName {
		t.Errorf("Expected CommonName %s, got %s", commonName, signedCertParsed.Subject.CommonName)
	}
	if !signedCertParsed.NotBefore.Before(time.Now()) || !signedCertParsed.NotBefore.Before(signedCertParsed.NotAfter) {
		t.Errorf("Invalid NotBefore field")
	}
	if signedCertParsed.KeyUsage != x509.KeyUsageDigitalSignature {
		t.Errorf("Invalid key usage")
	}
	if len(signedCertParsed.ExtKeyUsage) != 2 || signedCertParsed.ExtKeyUsage[0] != x509.ExtKeyUsageClientAuth || signedCertParsed.ExtKeyUsage[1] != x509.ExtKeyUsageServerAuth {
		t.Errorf("Invalid ExtKeyUsage")
	}

	// Verify the signed certificate with the CA certificate
	roots := x509.NewCertPool()
	caPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: caCert})
	ok := roots.AppendCertsFromPEM(caPem)
	if !ok {
		t.Fatalf("Failed to parse CA certificate")
	}

	opts := x509.VerifyOptions{
		Roots: roots,
	}

	_, err = signedCertParsed.Verify(opts)
	if err != nil {
		t.Fatalf("Fail to verify signed certificate: %v", err)
	}

	// Try to build tls.Certificate from signed certificate, CSR key, and CA certificate
	certPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: signedCert})
	keyPem := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key})

	_, err = tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		t.Fatalf("Failed to load signed certificate and key into tls.Certificate: %v", err)
	}
}
