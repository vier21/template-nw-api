package keys

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/vier21/pc-01-network-be/config"
)

func LoadPrivateKey() *rsa.PrivateKey {
	filePath := fmt.Sprintf("%s/%s", os.Getenv("APP_PATH"), config.GetConfig().RSAPrivPath)

	privKey, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Errorf("Cannot read private key: %s", err.Error())
		return nil
	}

	dec, _ := pem.Decode(privKey)

	rsaPriv, err := x509.ParsePKCS8PrivateKey(dec.Bytes)
	if err != nil {
		logrus.Errorf("Cannot parse private key: %s", err.Error())
		return nil
	}

	parsedKey := rsaPriv.(*rsa.PrivateKey)

	return parsedKey
}

func LoadPublicKey() *rsa.PublicKey {
	filePath := fmt.Sprintf("%s/%s", os.Getenv("APP_PATH"), config.GetConfig().RSAPubPath)

	pubKey, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	decpub, _ := pem.Decode(pubKey)
	rsapubd, err := x509.ParsePKIXPublicKey(decpub.Bytes)
	s := rsapubd.(*rsa.PublicKey)
	if err != nil {
		return nil
	}
	return s
}
