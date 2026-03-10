package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ============================================================
// WHAT IS A JWT?
// ============================================================
// JWT (JSON Web Token) is a compact, self-contained token used
// to securely transfer information between client and server.
//
// Common use cases:
//   - Authentication (login systems)
//   - Authorization (role-based access)
//   - Information exchange
//
// JWT is:
//   - Stateless  → server doesn't need to store sessions
//   - Signed     → tamper-proof via digital signature
//   - Compact    → easy to send in HTTP headers

// ============================================================
// STRUCTURE OF A JWT
// ============================================================
// A JWT looks like this:
//   eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMDEsInVzZXJuYW1lIjoia2FydGlrIn0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
//
// It has 3 parts separated by dots:
//   HEADER . PAYLOAD . SIGNATURE
//
// ── HEADER ──────────────────────────────────────────────────
// Contains metadata about the token.
//   {
//     "alg": "HS256",   ← signing algorithm (HMAC-SHA256)
//     "typ": "JWT"      ← token type
//   }
//
// ── PAYLOAD ─────────────────────────────────────────────────
// Contains claims (the actual data).
//   {
//     "user_id":  101,
//     "username": "kartik",
//     "exp":      1741542000   ← expiry timestamp
//   }
//
// Types of claims:
//   Registered → standard fields: exp, iat, iss, sub, aud
//   Public     → app data: user_id, email
//   Private    → custom fields: role, permissions
//
// ── SIGNATURE ───────────────────────────────────────────────
// Ensures the token hasn't been tampered with.
//   HMACSHA256(
//     base64(header) + "." + base64(payload),
//     secretKey
//   )
// If anyone modifies the payload → signature won't match → token rejected.

// ============================================================
// SETUP
// ============================================================
// Install the library:
//   go get github.com/golang-jwt/jwt/v5
//
// secretKey is used to sign and verify tokens.
// In production: load this from environment variables, never hardcode it.
//var secretKey = []byte("my_secret_key")

// ============================================================
// STEP 1: DEFINE CLAIMS
// ============================================================
// Claims are the data stored inside the token.
// We embed jwt.RegisteredClaims to get standard fields like exp, iat.
// type MyClaims struct {
// 	UserID   int    `json:"user_id"`
// 	Username string `json:"username"`
// 	jwt.RegisteredClaims
// }

// ============================================================
// STEP 2: GENERATE A TOKEN
// ============================================================
// generateToken creates and signs a JWT for a given user.
//
// Flow:
//  1. Build claims (payload data)
//  2. Create unsigned token with signing method
//  3. Sign the token with secretKey → get final JWT string
func generateToken(userID int, username string) (string, error) {
	// Build the claims (payload)
	claims := MyClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // issued right now
			Issuer:    "my-app",                                           // who issued the token
		},
	}

	// Create unsigned token
	// jwt.SigningMethodHS256 = HMAC with SHA-256 (symmetric, uses same key to sign & verify)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with secretKey → returns final JWT string
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return signedToken, nil
}

// ============================================================
// STEP 3: VERIFY A TOKEN
// ============================================================
// verifyToken parses and validates a JWT string.
//
// Flow:
//   1. Parse the token string
//   2. Validate signing method (prevent algorithm switching attacks)
//   3. Return claims if token is valid

// func verifyToken(tokenString string) (*MyClaims, error) {
// 	claims := &MyClaims{}

// 	// ParseWithClaims parses the token and fills claims
// 	// The keyfunc callback validates the signing method and returns the secret key
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

// 		// IMPORTANT: Always validate the signing method
// 		// This prevents "algorithm confusion attacks" where an attacker
// 		// sends a token signed with a different algorithm (e.g. "none")
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// Return the secret key for verification
// 		return secretKey, nil
// 	})
// 	if err != nil {
// 		// This catches: expired tokens, invalid signature, malformed tokens
// 		return nil, fmt.Errorf("error parsing token: %w", err)
// 	}

// 	if !token.Valid {
// 		return nil, fmt.Errorf("invalid token")
// 	}

// 	return claims, nil
// }

// ============================================================
// MAIN: SEE IT ALL IN ACTION
// ============================================================
// func main() {

// 	fmt.Println("=== JWT TOKEN DEMO ===")
// 	fmt.Println()

// 	// ── Generate ─────────────────────────────────────────────
// 	fmt.Println("1. Generating token for user kartik (ID: 1)...")
// 	tokenString, err := generateToken(1, "kartik")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Generated Token:")
// 	fmt.Println(tokenString)
// 	fmt.Println()

// 	// ── Verify valid token ───────────────────────────────────
// 	fmt.Println("2. Verifying the token...")
// 	claims, err := verifyToken(tokenString)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Token is VALID!")
// 	fmt.Println("  User ID  :", claims.UserID)
// 	fmt.Println("  Username :", claims.Username)
// 	fmt.Println("  Issued At:", claims.IssuedAt)
// 	fmt.Println("  Expires  :", claims.ExpiresAt)
// 	fmt.Println("  Issuer   :", claims.Issuer)
// 	fmt.Println()

// 	// ── Verify tampered token ────────────────────────────────
// 	fmt.Println("3. Trying a tampered token...")
// 	tamperedToken := tokenString + "tampered"
// 	_, err = verifyToken(tamperedToken)
// 	if err != nil {
// 		fmt.Println("Correctly rejected tampered token:", err)
// 	}
// 	fmt.Println()

// 	// ── Verify fake token ────────────────────────────────────
// 	fmt.Println("4. Trying a completely fake token...")
// 	fakeToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5OTksInVzZXJuYW1lIjoiaGFja2VyIn0.fakesignature"
// 	_, err = verifyToken(fakeToken)
// 	if err != nil {
// 		fmt.Println("Correctly rejected fake token:", err)
// 	}
// }
// ```

// **Output you'll see:**
// ```
// === JWT TOKEN DEMO ===

// 1. Generating token for user kartik (ID: 1)...
// Generated Token:
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6Imthc...

// 2. Verifying the token...
// Token is VALID!
//   User ID  : 1
//   Username : kartik
//   Issued At: 2024-01-01 10:00:00 +0000 UTC
//   Expires  : 2024-01-02 10:00:00 +0000 UTC
//   Issuer   : my-app

// 3. Trying a tampered token...
// Correctly rejected tampered token: error parsing token: ...

// 4. Trying a completely fake token...
// Correctly rejected fake token: error parsing token: ...
