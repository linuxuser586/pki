package cert

import (
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"
)

const testKey = "test key"
const testCert = "test cert"

func TestDevPublicKey(t *testing.T) {
	c := CAPublic()
	if c != caCert {
		t.Errorf("got: %v, want: %v", c, caCert)
	}
}

func TestEnvPublicKey(t *testing.T) {
	os.Setenv(keyEnv, testKey)
	os.Setenv(certEnv, testCert)
	setup()
	c := CAPublic()
	if caKey != testKey {
		t.Errorf("test key got: %v, want: %v", c, testKey)
	}
	if c != testCert {
		t.Errorf("test cert got: %v, want: %v", c, testCert)
	}
	os.Setenv(keyEnv, "")
	os.Setenv(certEnv, "")
	setup()
}

func TestValidCert(t *testing.T) {
	c, err := GenerateRequest([]string{"test.example.com", "172.17.0.2"})
	if err != nil {
		t.Fatalf("got error %v", err)
	}
	p, _ := pem.Decode([]byte(c.Cert))
	caCrt, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		t.Fatalf("got error %v", err)
	}
	l := len(caCrt.DNSNames)
	expected := 1
	if l != expected {
		t.Errorf("DNS count got: %v %v, want: %v", l, caCrt.DNSNames, expected)
	}
	l = len(caCrt.IPAddresses)
	expected = 2
	if l != expected {
		t.Errorf("IP count got: %v %v, want: %v", l, caCrt.IPAddresses, expected)
	}
}
