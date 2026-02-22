package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var signingMethods = map[string]jwt.SigningMethod{
	"HS256": jwt.SigningMethodHS256,
	"HS384": jwt.SigningMethodHS384,
	"HS512": jwt.SigningMethodHS512,
	"RS256": jwt.SigningMethodRS256,
	"RS384": jwt.SigningMethodRS384,
	"RS512": jwt.SigningMethodRS512,
}

type JWTDecodeResult struct {
	Token            *jwt.Token
	Error            error
	IsTokenValid     bool
	IsSignatureValid bool
}

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

func (r *JWTDecodeResult) Valid() bool {
	return r.Error == nil && r.IsTokenValid && r.IsSignatureValid
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

		if result.IsTokenValid && result.IsSignatureValid {
			result.Error = err
		}
	}

	return &result
}

func ParseECDSAPublicKeyFromPEM(pemBytes []byte) (any, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing public key")
	}

	return x509.ParsePKIXPublicKey(block.Bytes)
}

type JWTEncodeResult struct {
	Token        string
	HeaderError  string
	PayloadError string
	SigningError string
}

func JWTEncodeToken(header map[string]interface{}, claims jwt.MapClaims, secret string) *JWTEncodeResult {
	result := &JWTEncodeResult{}

	var signingMethod jwt.SigningMethod = jwt.SigningMethodHS256
	if alg, ok := header["alg"]; ok {
		if algStr, ok := alg.(string); ok {
			if method, exists := signingMethods[algStr]; exists {
				signingMethod = method
			}
		}
	}

	token := jwt.NewWithClaims(signingMethod, claims)

	for k, v := range header {
		if k != "alg" && k != "typ" { // Don't override alg and typ
			token.Header[k] = v
		}
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		result.SigningError = "Error signing token: " + err.Error()
		return result
	}

	result.Token = tokenString
	return result
}
