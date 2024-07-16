package utils

import (
	"fmt"
	"os"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenstring string) (*jwt.Token, error) {
	WorkOSJWKSURL := fmt.Sprintf("https://api.workos.com/sso/jwks/%s", os.Getenv("WORKOS_CLIENT_ID"))
	//fmt.Println(WorkOSJWKSURL)

	jwks, err := keyfunc.NewDefault([]string{WorkOSJWKSURL})
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
