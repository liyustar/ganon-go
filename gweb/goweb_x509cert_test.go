package gweb_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"testing"
	"time"
)

func TestGoWebX509Cert(t *testing.T) {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		CommonName:   "Go Web Programming",
		Organization: []string{"Li Yuxing Co."},
	}
	now := time.Now()
	then := now.Add(365 * 24 * time.Hour) // one year
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: subject,
		//	NotBefore: time.Unix(now, 0).UTC(),
		//	NotAfter:  time.Unix(now+60*60*24*365, 0).UTC(),
		NotBefore: now,
		NotAfter:  then,

		//SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,

		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},

		//BasicConstraintsValid: true,
		//IsCA:                  true,
		//DNSNames:              []string{"jan.newmarch.name", "localhost"},
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	derBytes, err := x509.CreateCertificate(rand.Reader, &template,
		&template, &pk.PublicKey, pk)
	checkError(err)

	certPEMFile, err := os.Create("cert.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certPEMFile.Close()

	keyPEMFile, err := os.Create("private.pem")
	checkError(err)
	pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyPEMFile.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
