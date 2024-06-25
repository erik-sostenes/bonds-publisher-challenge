package logic

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"log/slog"
	"time"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
	gojwt "github.com/golang-jwt/jwt/v5"
)

type (
	tokenGenerator struct {
		privateKey *rsa.PrivateKey
	}

	tokenValidator struct {
		publicKey *rsa.PublicKey
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
	authorization := &domain.Authorization{
		UserID:   user.ID(),
		UserName: user.Name(),
		Role: struct {
			Id          uint8
			RoleType    string
			Permissions uint8
		}{
			Id:          user.Role().ID(),
			RoleType:    user.Role().Type(),
			Permissions: permissions,
		},
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

func NewTokenValidator(publicKeyBase64 string) ports.TokenValidator {
	publicKeyPEM, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		panic(err)
	}

	publicKey, err := gojwt.ParseRSAPublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		panic(err)
	}

	return &tokenValidator{
		publicKey: publicKey,
	}
}

// Validate method that validates the token using the public key and the type of encryption method
func (jwt *tokenValidator) Validate(strToken string) (*domain.Authorization, error) {
	token, err := gojwt.Parse(strToken, func(t *gojwt.Token) (any, error) {
		return jwt.publicKey, nil
	})
	if err != nil {
		switch err {
		case gojwt.ErrTokenMalformed:
			return nil, errors.New("the provided JWT token is malformed")
		case gojwt.ErrTokenSignatureInvalid:
			return nil, errors.New("the signature of the JWT token is invalid")
		case gojwt.ErrTokenExpired:
			return nil, errors.New("the JWT token has expired")
		default:
			return nil, err
		}
	}

	if claims, ok := token.Claims.(gojwt.MapClaims); ok && token.Valid {
		auth := &domain.Authorization{
			UserID:   claims["UserID"].(string),
			UserName: claims["UserName"].(string),
			Role: struct {
				Id          uint8
				RoleType    string
				Permissions uint8
			}{
				Id:          uint8(claims["Role"].(map[string]any)["Id"].(float64)),
				RoleType:    claims["Role"].(map[string]any)["RoleType"].(string),
				Permissions: uint8(claims["Role"].(map[string]any)["Permissions"].(float64)),
			},
		}
		return auth, nil
	}
	return nil, nil
}
