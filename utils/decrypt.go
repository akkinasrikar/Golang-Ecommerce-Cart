package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func parsePrivateKeyFromCertificate(key string) (*rsa.PrivateKey, error) {
	cer, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(cer)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the private key")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("failed to parse PEM block containing the private key")
	}

	return rsaPrivateKey, nil
}

func DecryptData(data []byte, key string) ([]byte, error) {
	privateKey, err := parsePrivateKeyFromCertificate(key)
	if err != nil {
		return nil, err
	}
	decodedBase64Data, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	hash := sha512.New()
	decryptedData, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, decodedBase64Data, nil)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil
}
