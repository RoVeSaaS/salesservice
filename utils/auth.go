package utils

import (
	"fmt"
	"os"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

func VerifyToken(tokenstring string) (*jwt.Token, error) {
	usermanagement.SetAPIKey(os.Getenv("WORKOS_API_KEY"))
	WorkOSJWKSURL, err := usermanagement.GetJWKSURL(os.Getenv("WORKOS_CLIENT_ID"))
	if err != nil {
		return nil, err
	}
	//fmt.Println(WorkOSJWKSURL)

	jwks, err := keyfunc.NewDefault([]string{WorkOSJWKSURL.String()})
	if err != nil {
		panic(err)
	}
	token, err := jwt.Parse(tokenstring, jwks.Keyfunc)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
