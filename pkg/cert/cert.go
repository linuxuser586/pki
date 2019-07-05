package cert

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"net"
	"os"
	"time"

	"crypto/rand"
)

const devkey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIAhKKOC08IykgXF4RuCj8SAYcqnOp8znUIIA0I7+a96toAoGCCqGSM49
AwEHoUQDQgAEO2B8nO1LFDZSEi1HGWTl4v8sDWgUkET/9msP5teECW5IWs8CoVUQ
sTGY1hf9Vbkpgwm0irIlhk9MxmJ/6cK/zw==
-----END EC PRIVATE KEY-----`

const devcert = `-----BEGIN CERTIFICATE-----
MIIBkDCCATagAwIBAgIQGfNC73kPKFoun3DpJiz7VDAKBggqhkjOPQQDAjAoMSYw
JAYDVQQDEx1Qb2x5YXBwaWMgRGV2ZWxvcG1lbnQgUm9vdCBDQTAeFw0xODA5MjAx
MjQ5MzZaFw0yMzA5MTkxMjQ5MzZaMCgxJjAkBgNVBAMTHVBvbHlhcHBpYyBEZXZl
bG9wbWVudCBSb290IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEO2B8nO1L
FDZSEi1HGWTl4v8sDWgUkET/9msP5teECW5IWs8CoVUQsTGY1hf9Vbkpgwm0irIl
hk9MxmJ/6cK/z6NCMEAwDgYDVR0PAQH/BAQDAgKkMB0GA1UdJQQWMBQGCCsGAQUF
BwMBBggrBgEFBQcDAjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUC
IEV2LiE+TAt8aF4oS6ga52lXAyt/bPa4Dyxj8HpO3sFEAiEAp/iZ+aF+gW7PoTox
4nhqA2kPV47TtWeQozd9LTduIP8=
-----END CERTIFICATE-----`

const gracePeriod = -5

// 400 days
const hours = 9600

var errInvalidDNS = errors.New("DNS must have at lest one entry")

const keyEnv = "PKI_CA_KEY"
const certEnv = "PKI_CA_CERT"

var caKey string
var caCert string

func init() {
	setup()
}

func setup() {
	caKey = os.Getenv(keyEnv)
	caCert = os.Getenv(certEnv)
	if caKey == "" || caCert == "" {
		log.Warn("using development ca key/cert")
		caKey = devkey
		caCert = devcert
	}
}

func toPem(k *ecdsa.PrivateKey) (string, error) {
	b, err := x509.MarshalECPrivateKey(k)
	if err != nil {
		return "", err
	}
	priv := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: b,
	}
	o := new(bytes.Buffer)
	if err := pem.Encode(o, priv); err != nil {
		return "", err
	}
	return o.String(), nil
}

// Client holds the client request private key and public certificate
type Client struct {
	Key  string
	Cert string
}

// GenerateRequest generates a new client key/cert pair
func GenerateRequest(DNSNames []string) (*Client, error) {
	if len(DNSNames) < 1 {
		return nil, errInvalidDNS
	}
	k, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	p, _ := pem.Decode([]byte(caKey))
	caPriv, err := x509.ParseECPrivateKey(p.Bytes)
	if err != nil {
		return nil, err
	}
	p, _ = pem.Decode([]byte(caCert))
	caCrt, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		return nil, err
	}
	notBefore := time.Now().Add(time.Duration(-5) * time.Minute)
	notAfter := notBefore.Add(time.Duration(9600) * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, err
	}

	cn := DNSNames[0]
	var dns []string
	ip := []net.IP{[]byte{127, 0, 0, 1}}
	for _, s := range DNSNames {
		p := net.ParseIP(s)
		if p != nil {
			ip = append(ip, p)
		} else {
			dns = append(dns, s)
		}
	}

	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: cn,
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageKeyAgreement,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		DNSNames:              dns,
		IPAddresses:           ip,
		SignatureAlgorithm:    x509.ECDSAWithSHA256,
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, template, caCrt, &k.PublicKey, caPriv)
	if err != nil {
		return nil, err
	}
	crt := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}
	o := new(bytes.Buffer)
	if err := pem.Encode(o, &crt); err != nil {
		return nil, err
	}
	pem, err := toPem(k)
	if err != nil {
		return nil, err
	}
	req := &Client{Key: pem, Cert: o.String()}
	return req, nil
}

// CAPublic provides the certificate authority public certificate
func CAPublic() string {
	return caCert
}
