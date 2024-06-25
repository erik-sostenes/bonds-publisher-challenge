package domain

import (
	gojwt "github.com/golang-jwt/jwt/v5"
)

type (
	// RegisteredClaims is a structure representing JWT own claim types
	RegisteredClaims = gojwt.RegisteredClaims

	// Authorization is a structure representing the payload of a token
	Authorization struct {
		UserID,
		UserName string
		Role struct {
			Id          uint8
			RoleType    string
			Permissions uint8
		}
		RegisteredClaims
	}
)
