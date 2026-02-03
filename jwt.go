package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// JWTDecodeResult holds the result of JWT decoding operation
type JWTDecodeResult struct {
	Token            *jwt.Token
	Error            error
	IsTokenValid     bool
	IsSignatureValid bool
}

// JsonMarshaledHeader returns the JSON marshaled header of the JWT token
func (r *JWTDecodeResult) JsonMarshaledHeader() string {
	if r.Token == nil {
		return ""
	}

	v, err := json.MarshalIndent(r.Token.Header, "", "  ")
	if err != nil {
		return ""
	}

	return string(v)
}

// JsonMarshaledClaims returns the JSON marshaled claims of the JWT token
func (r *JWTDecodeResult) JsonMarshaledClaims() string {
	if r.Token == nil {
		return ""
	}

	v, err := json.MarshalIndent(r.Token.Claims, "", "  ")
	if err != nil {
		return ""
	}

	return string(v)
}

// Valid returns whether the JWT is valid based on token validity and signature verification
func (r *JWTDecodeResult) Valid() bool {
	return r.Error == nil && r.IsTokenValid && r.IsSignatureValid
}

// JWTDecodeToken decodes and validates a JWT token using the provided secret
func JWTDecodeToken(token, secret string) *JWTDecodeResult {
	parsedToken, err := jwt.Parse(token, jwt.Keyfunc(func(t *jwt.Token) (any, error) {
		pubKey, err := ParseECDSAPublicKeyFromPEM([]byte(secret))
		if err != nil {
			return []byte(secret), nil
		}
		return pubKey, nil
	}))

	result := JWTDecodeResult{
		Token:            parsedToken,
		IsTokenValid:     true,
		IsSignatureValid: true,
	}

	if err != nil {
		result.IsTokenValid = !strings.Contains(err.Error(), "token is malformed")
		result.IsSignatureValid = !strings.Contains(err.Error(), "token signature is invalid") && result.IsTokenValid

		// Unexpected error occurred
		if result.IsTokenValid && result.IsSignatureValid {
			result.Error = err
		}
	}

	return &result
}

// ParseECDSAPublicKeyFromPEM attempts to parse an ECDSA public key from PEM format
func ParseECDSAPublicKeyFromPEM(pemBytes []byte) (any, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing public key")
	}

	return x509.ParsePKIXPublicKey(block.Bytes)
}
