package logic

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"log/slog"
	"time"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
	gojwt "github.com/golang-jwt/jwt/v5"
)

type (
	// RegisteredClaims is a structure representing JWT own claim types
	RegisteredClaims = gojwt.RegisteredClaims

	// Authorization is a structure representing the payload of a token
	Authorization struct {
		UserID,
		UserName string
		Role        map[string]any
		Permissions uint
		RegisteredClaims
	}

	tokenGenerator struct {
		privateKey *rsa.PrivateKey
	}
)

func NewTokenGenerator(privateKeyBase64 string) ports.TokenGenerator {
	if privateKeyBase64 == "" {
		panic(`missing privateKeyBase64`)
	}

	privateKeyPEM, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		panic(err)
	}

	privateKey, err := gojwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	return &tokenGenerator{
		privateKey: privateKey,
	}
}

// Generate a new [Token] with the specified signing method and claims
func (jwt tokenGenerator) Generate(user *domain.User, permissions uint8) (token string, err error) {
	role := map[string]any{
		"id":   user.Role().ID(),
		"type": user.Role().Type(),
	}

	authorization := &Authorization{
		UserID:      user.ID(),
		UserName:    user.Name(),
		Role:        role,
		Permissions: uint(permissions),
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			Issuer:    "cicada",
		},
	}

	jwtToken := gojwt.NewWithClaims(gojwt.SigningMethodRS256, authorization)

	tokenSigned, err := jwtToken.SignedString(jwt.privateKey)
	if err != nil {
		slog.ErrorContext(context.Background(), "error generating token", "msg", err)
		return
	}

	return tokenSigned, nil
}
