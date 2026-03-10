package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// A JWT is a secure token used to transfer information between client and server.
// JWT (JSON Web Token) is widely used in authentication and authorization systems in backend applications.
// JWT is: Stateless, Secure (digitally signed), Compact, Used for authentication

// Structure of JWT: A JWT has 3 parts. -> HEADER.PAYLOAD.SIGNATURE

// Header -> Contains token metadata.
// {
//  "alg": "HS256",
//  "typ": "JWT"
// }
// | Field | Meaning        |
// | ----- | -------------- |
// | alg   | algorithm used |
// | typ   | token type     |

// Payload: Contains claims (data).
// {
//  "user_id": 101,
//  "username": "kartik",
//  "role": "admin",
//  "exp": 1741542000
// }

// Types of claims:
// Registered claims
// | Claim | Meaning    |
// | ----- | ---------- |
// | exp   | expiration |
// | iat   | issued at  |
// | iss   | issuer     |
// | sub   | subject    |

// Public claims: Application data -> user_id, email,
// Private claims: Custom application fields -> ticket_category, permissions

// Signature: Used to verify token integrity -> If someone modifies payload → signature fails.

// JWT Library in Go: github.com/golang-jwt/jwt/v5
// Install: go get github.com/golang-jwt/jwt/v5

// Creating JWT in Go
var secretKey = []byte("my_secret_key")

type MyClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateJwtTokens(userID int, username string) (string, error) {
	claims := MyClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// func verifyToken(tokenString string) (*MyClaims, error) {
// 	claims := &MyClaims{}

// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return secretKey, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, fmt.Errorf("invalid token")
// 	}

// 	return claims, nil
// }

// func jwtTokens() {
// 	// Generate token
// 	token, err := generateJwtTokens(1, "kartik")
// 	if err != nil {
// 		fmt.Println("Error generating token:", err)
// 		return
// 	}
// 	fmt.Println("JWT Token:", token)

// 	// Verify token
// 	claims, err := verifyToken(token)
// 	if err != nil {
// 		fmt.Println("Error verifying token:", err)
// 		return
// 	}
// 	fmt.Println("Token valid!")
// 	fmt.Println("User ID:", claims.UserID)
// 	fmt.Println("Username:", claims.Username)
// 	fmt.Println("Expires At:", claims.ExpiresAt)
// }

// func main() {
// 	jwtTokens()
// }
