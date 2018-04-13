package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"golang.org/x/crypto/pkcs12"
)

// Certificate errors
var (
	ErrExpired = errors.New("certificate has expired or is not yet valid")
)

func main() {
	b, err := ioutil.ReadFile("eduroam.p12")
	if err != nil {
		log.Fatal(err)
	}
	password, err := ioutil.ReadFile("passwd.txt")
	if err != nil {
		log.Fatal(err)
	}
	password = bytes.TrimSuffix(password, []byte("\n"))

	switch "use-ext" {
	case "use-go":
		cert, pkey, err := decodeGo(b, password)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("cert: %v", cert)
		log.Printf("pkey: %v", pkey)
	case "use-ext":
		err := decodeOpenSSL(b, password)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func decodeOpenSSL(p12, passwd []byte) error {
	{
		cmd := exec.Command(
			"openssl", "pkcs12",
			"-nocerts",
			"-in", "eduroam.p12",
			"-out", "userkey.pem",
			"-passin", "file:passwd.txt",
			"-passout", "file:./passwd.txt",
		)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("could not generate user-key: %v", err)
		}
	}
	{
		cmd := exec.Command(
			"openssl", "pkcs12",
			"-clcerts", "-nokeys",
			"-in", "eduroam.p12",
			"-out", "usercert.pem",
			"-passin", "file:passwd.txt",
			"-passout", "file:./passwd.txt",
		)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("could not generate user-cert: %v", err)
		}
	}

	return nil
}

// decode and verify an in memory .p12 certificate (DER binary format).
func decodeGo(p12 []byte, password []byte) (*x509.Certificate, *rsa.PrivateKey, error) {
	blocks, err := pkcs12.ToPEM(p12, string(password))
	if err != nil {
		log.Fatalf("to-PEM: %v", err)
	}
	log.Printf(">>> blocks: %v", blocks)
	// decode an x509.Certificate to verify
	pkey, cert, err := pkcs12.Decode(p12, string(password))
	if err != nil {
		return nil, nil, err
	}
	if err := verify(cert); err != nil {
		return nil, nil, err
	}

	// assert that private key is RSA
	priv, ok := pkey.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, errors.New("expected RSA private key type")
	}
	return cert, priv, nil
}

// verify checks if a certificate has expired
func verify(cert *x509.Certificate) error {
	_, err := cert.Verify(x509.VerifyOptions{})
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case x509.CertificateInvalidError:
		switch e.Reason {
		case x509.Expired:
			return ErrExpired
		default:
			return err
		}
	case x509.UnknownAuthorityError:
		// Apple cert isn't in the cert pool
		// ignoring this error
		return nil
	default:
		return err
	}
}
