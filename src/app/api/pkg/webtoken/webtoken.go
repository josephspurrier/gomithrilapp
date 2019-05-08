package webtoken

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// ErrBadTokenFormat is when the refresh token has expired.
	ErrBadTokenFormat = errors.New("authorization header format must be Bearer {token}")
	// ErrMissingExpiration is when the token is missing an expiration date.
	ErrMissingExpiration = errors.New("token must contain an expiration")
	// ErrNoToken is when the refresh token has expired.
	ErrNoToken = errors.New("no access token provided")
	// ErrBadToken is when the refresh token has expired.
	ErrBadToken = errors.New("bad token")
	// ErrBadClock is when the clock is not initialized.
	ErrBadClock = errors.New("bad clock")
	// ErrBadSecret is when the secret is left nil or blank which is unsecure.
	ErrBadSecret = errors.New("secret cannot be nil or blank")
)

// CustomClaims is the payload for the JWT.
type CustomClaims struct {
	jwt.StandardClaims
	UserID string `json:"userID"`
}

// JWTTokenExtractor extracts tokens from http header
type JWTTokenExtractor struct {
}

// GetToken gets token out of http header
func (e *JWTTokenExtractor) GetToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", ErrNoToken
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", ErrBadTokenFormat
	}

	return authHeaderParts[1], nil
}

// Clock is a real clock.
type Clock struct{}

// Now returns the current time.
func (Clock) Now() time.Time {
	return time.Now()
}

// After returns the time after.
/*func (Clock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}*/

// IClock represents a real clock.
type IClock interface {
	Now() time.Time
	//After(d time.Duration) <-chan time.Time
}

// JWTAuth is the struct that verifies and generates jwt access tokens
type JWTAuth struct {
	Clock                IClock
	PrivateKey           *[]byte
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

// AuthToken represents a JWT.
type AuthToken struct {
	Token  string
	Expiry time.Time
}

// User represents a person.
type User struct {
	ID string
}

// Verify verifies a token is valid.
func (p *JWTAuth) Verify(token string) (string, error) {
	// Prevent crashing from the clock.
	if p.Clock == nil {
		return "", ErrBadClock
	}

	// Don't allow an empty secret.
	if p.PrivateKey == nil || *p.PrivateKey == nil || len(*p.PrivateKey) == 0 {
		return "", ErrBadSecret
	}

	// Sync the clock of the package.
	jwt.TimeFunc = func() time.Time {
		return p.Clock.Now()
	}

	// Create a map to store our claims.
	accessClaims := new(CustomClaims)

	// Parse the JWT.
	t, err := jwt.ParseWithClaims(token, accessClaims, func(token *jwt.Token) (interface{}, error) {
		return *p.PrivateKey, nil
	})
	// The parse with claims will always return an error, but won't always be
	// valid because it could be null.
	if err != nil || !t.Valid {
		return "", err
	}

	// Require an expiration date.
	if accessClaims.ExpiresAt == 0 {
		return "", ErrMissingExpiration
	}

	// Return the user ID.
	//return fmt.Sprint(accessClaims["UserID"]), nil
	return fmt.Sprint(accessClaims.UserID), nil
}

// RefreshAccessToken generates a new access token off the old access token.
func (p *JWTAuth) RefreshAccessToken(oldToken string) (accessToken *AuthToken, err error) {
	// Prevent crashing from the clock.
	if p.Clock == nil {
		return nil, ErrBadClock
	}

	// Don't allow an empty secret.
	if p.PrivateKey == nil || *p.PrivateKey == nil || len(*p.PrivateKey) == 0 {
		return nil, ErrBadSecret
	}

	// Sync the clock of the package.
	jwt.TimeFunc = func() time.Time {
		return p.Clock.Now()
	}

	// Create a map to store our claims.
	accessClaims := new(CustomClaims)

	// Load old claims.
	t, err := jwt.ParseWithClaims(oldToken, accessClaims, func(token *jwt.Token) (interface{}, error) {
		return *p.PrivateKey, nil
	})
	// The parse with claims will always return an error, but won't always be
	// valid because it could be null.
	if err != nil || !t.Valid {
		return nil, err
	}

	// Require an expiration date.
	if accessClaims.ExpiresAt == 0 {
		return nil, ErrMissingExpiration
	}

	// Update the token expiration.
	accessTokenExpiry := p.Clock.Now().Add(p.AccessTokenDuration)
	accessClaims.ExpiresAt = accessTokenExpiry.Unix()

	// Create the token.
	accToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	// Sign the token with our secret.
	aToken, err := accToken.SignedString(*p.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &AuthToken{Token: aToken, Expiry: accessTokenExpiry}, nil
}

// GenerateTokens will generate the access token and refresh token.
func (p *JWTAuth) GenerateTokens(user *User) (accessToken *AuthToken, refreshToken *AuthToken, err error) {
	// Prevent crashing from the clock.
	if p.Clock == nil {
		return nil, nil, ErrBadClock
	}

	// Don't allow an empty secret.
	if p.PrivateKey == nil || *p.PrivateKey == nil || len(*p.PrivateKey) == 0 {
		return nil, nil, ErrBadSecret
	}

	// This is not required for this package. It never gets called so it
	// is never tested.
	// Sync the clock of the package.
	/*jwt.TimeFunc = func() time.Time {
		return p.Clock.Now()
	}*/

	// Create a clock to use.
	clock := p.Clock.Now()

	// Create a map to store our claims.
	accessClaims := CustomClaims{}
	accessClaims.UserID = user.ID
	accessTokenExpiry := clock.Add(p.AccessTokenDuration)
	accessClaims.ExpiresAt = accessTokenExpiry.Unix()

	// Create the token.
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	// Sign the token with our secret.
	accToken, err := aToken.SignedString(*p.PrivateKey)
	if err != nil {
		return nil, nil, err
	}

	// Create a map to store our claims.
	refreshClaims := CustomClaims{}
	refreshClaims.UserID = user.ID
	refreshTokenExpiry := clock.Add(p.RefreshTokenDuration)
	refreshClaims.ExpiresAt = refreshTokenExpiry.Unix()

	// Create the token.
	rToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Sign the token with our secret.
	refToken, err := rToken.SignedString(*p.PrivateKey)
	if err != nil {
		return nil, nil, err
	}

	return &AuthToken{Token: accToken, Expiry: accessTokenExpiry},
		&AuthToken{Token: refToken, Expiry: refreshTokenExpiry}, nil
}
