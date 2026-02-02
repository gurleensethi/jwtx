package main

import (
	"encoding/json"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JWTDecodeResult struct {
	Token              *jwt.Token
	Error              error
	IsTokenInvalid     bool
	IsSignatureInvalid bool
}

func (r JWTDecodeResult) JsonMarshedHeader() string {
	if r.Token == nil {
		return ""
	}

	v, _ := json.MarshalIndent(r.Token.Header, "", "  ")

	return string(v)
}

func (r JWTDecodeResult) JsonMarshledClaims() string {
	if r.Token == nil {
		return ""
	}

	v, _ := json.MarshalIndent(r.Token.Claims, "", "  ")

	return string(v)
}

func JWTDecodeToken(token, secret string) JWTDecodeResult {
	parsedToken, err := jwt.Parse(token, jwt.Keyfunc(func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	}))

	result := JWTDecodeResult{
		Token: parsedToken,
		Error: err,
	}

	if err != nil {
		if strings.Contains(err.Error(), "token is malformed") {
			result.IsTokenInvalid = true
		}

		if strings.Contains(err.Error(), "token signature is invalid") {
			result.IsSignatureInvalid = true
		}
	}

	return result
}
