package main

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
)

type JWTDecodeResult struct {
	Token *jwt.Token
	Error error
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

	return JWTDecodeResult{
		Token: parsedToken,
		Error: err,
	}
}
