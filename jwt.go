package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JWTDecodeResult struct {
	Token            *jwt.Token
	Error            error
	IsTokenValid     bool
	IsSignatureValid bool
}

func (r *JWTDecodeResult) JsonMarshedHeader() string {
	if r.Token == nil {
		return ""
	}

	v, _ := json.MarshalIndent(r.Token.Header, "", "  ")

	return string(v)
}

func (r *JWTDecodeResult) JsonMarshledClaims() string {
	if r.Token == nil {
		return ""
	}

	v, _ := json.MarshalIndent(r.Token.Claims, "", "  ")

	return string(v)
}

func (r *JWTDecodeResult) Valid() bool {
	return r.Error != nil && r.IsTokenValid && r.IsSignatureValid
}

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

func ParseECDSAPublicKeyFromPEM(pemBytes []byte) (any, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block contaning public key")
	}

	return x509.ParsePKIXPublicKey(block.Bytes)
}
