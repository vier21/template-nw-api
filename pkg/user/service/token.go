package service

import (
	"context"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/vier21/pc-01-network-be/pkg/user/domain"
)

type IToken interface {
	GenerateToken(context.Context, domain.User) (string, *jwt.Token, error)
}

type JWTClaims struct {
	DataUserAuthenticated
	jwt.RegisteredClaims
}

type tokenMaker struct {
	privKey *rsa.PrivateKey
	pubKey  *rsa.PublicKey
}

func NewTokenRSA(priv *rsa.PrivateKey, pub *rsa.PublicKey) IToken {
	return &tokenMaker{
		privKey: priv,
		pubKey:  pub,
	}
}

func (t *tokenMaker) GenerateToken(ctx context.Context, user domain.User) (string, *jwt.Token, error) {
	claims := &JWTClaims{
		*setAuthenticatedUser(&user),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(1*time.Hour).Unix(), 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        "1",
			Issuer:    "github.com/vier21/pc-01",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(t.privKey)
	if err != nil {
		logrus.Errorf("jwt sign error: %s \n", err.Error())
		return "", nil, err
	}
	return ss, token, nil
}
