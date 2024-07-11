package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(tokenstring string) (*jwt.Token, error) {
	WorkOSJWKSURL := fmt.Sprintf("https://api.workos.com/sso/jwks/%s", os.Getenv("WORKOS_CLIENT_ID"))
	//fmt.Println(WorkOSJWKSURL)
	options := keyfunc.Options{
		RefreshInterval: time.Hour, // Refresh JWKS periodically
		RefreshErrorHandler: func(err error) {
			// Handle errors if any during refresh
			log.Printf("Error refreshing JWKS: %v", err)
		},
	}
	jwks, err := keyfunc.Get(WorkOSJWKSURL, options)
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
