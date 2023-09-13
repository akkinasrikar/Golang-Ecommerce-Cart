package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

func loadPublicCertificate() (string, error) {
	// get root path
	publicCertificate, err := os.ReadFile("certificates/public.cer")
	if err != nil {
		return "", err
	}
	return string(publicCertificate), nil
}

func parsePublicKeyFromCertificate(cer string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(cer))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return publicKey.(*rsa.PublicKey), nil
}

func EncryptData(data []byte) (string, error) {
	publicCertificate, err := loadPublicCertificate()
	if err != nil {
		return "", err
	}
	publicKey, err := parsePublicKeyFromCertificate(publicCertificate)
	if err != nil {
		return "", err
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}
