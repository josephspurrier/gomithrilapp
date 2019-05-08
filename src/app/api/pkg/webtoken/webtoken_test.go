package webtoken_test

import (
	"testing"
	"time"

	"app/api/pkg/webtoken"

	"github.com/stretchr/testify/assert"
)

func TestNoToken(t *testing.T) {
	j := webtoken.JWTTokenExtractor{}

	s, err := j.GetToken("")
	assert.Equal(t, webtoken.ErrNoToken, err)
	assert.Equal(t, "", s)
}

func TestBadTokenFormat(t *testing.T) {
	j := webtoken.JWTTokenExtractor{}

	s, err := j.GetToken("bad")
	assert.Equal(t, webtoken.ErrBadTokenFormat, err)
	assert.Equal(t, "", s)
}

func TestGoodToken(t *testing.T) {
	j := webtoken.JWTTokenExtractor{}

	s, err := j.GetToken("Bearer foo")
	assert.Equal(t, nil, err)
	assert.Equal(t, "foo", s)
}

func TestBadClock(t *testing.T) {
	b := []byte("s3cr3tp@$$w0rdk33")

	j := &webtoken.JWTAuth{
		Clock:                nil,
		PrivateKey:           &b,
		RefreshTokenDuration: time.Hour * 24 * 7,
		AccessTokenDuration:  time.Hour * 2,
	}

	u := new(webtoken.User)
	u.ID = "100"

	a, err := j.RefreshAccessToken("")
	assert.NotEqual(t, nil, a)
	assert.Equal(t, webtoken.ErrBadClock, err)

	a, r, err := j.GenerateTokens(u)
	assert.NotEqual(t, nil, a)
	assert.NotEqual(t, nil, r)
	assert.Equal(t, webtoken.ErrBadClock, err)

	s, err := j.Verify("")
	assert.NotEqual(t, nil, s)
	assert.Equal(t, webtoken.ErrBadClock, err)
}

func TestGoodGenerateVerifyToken(t *testing.T) {
	b := []byte("Pa$$w0rd")

	j := &webtoken.JWTAuth{
		Clock:                webtoken.Clock{},
		PrivateKey:           &b,
		RefreshTokenDuration: time.Hour * 24 * 7,
		AccessTokenDuration:  time.Hour * 2,
	}

	u := new(webtoken.User)
	u.ID = "100"

	a, r, err := j.GenerateTokens(u)

	assert.NotNil(t, a)
	assert.NotNil(t, r)
	assert.Equal(t, nil, err)

	s, err := j.Verify(a.Token)
	assert.NotNil(t, s)
	t.Log(err)
	assert.Equal(t, nil, err)
}

func TestBadGenerateVerifyToken(t *testing.T) {
	b := []byte("Pa$$w0rd")

	j := &webtoken.JWTAuth{
		Clock:                webtoken.Clock{},
		PrivateKey:           &b,
		RefreshTokenDuration: time.Hour * 24 * 7,
		AccessTokenDuration:  time.Hour * 2,
	}

	u := new(webtoken.User)
	u.ID = "100"

	a, r, err := j.GenerateTokens(u)

	assert.NotNil(t, a)
	assert.NotNil(t, r)
	assert.Equal(t, nil, err)

	s, err := j.Verify("bad")
	assert.Equal(t, "", s)
	assert.NotEqual(t, nil, err)

	s, err = j.Verify("")
	assert.Equal(t, "", s)
	assert.NotEqual(t, nil, err)
}

func TestBadGenerateVerifyTokenReal(t *testing.T) {
	b := []byte("Pa$$w0rd")

	j := &webtoken.JWTAuth{
		Clock:                webtoken.Clock{},
		PrivateKey:           &b,
		RefreshTokenDuration: time.Hour * 24 * 7,
		AccessTokenDuration:  time.Hour * 2,
	}

	s, err := j.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsInVzZXJJRCI6Mn0.rN-m8KD3oLK_-kW3vKaialQ06Fy_k_d95cPQJyalTJE")
	assert.Equal(t, "", s)
	assert.NotEqual(t, nil, err)
}

func TestNotValidGenerateVerifyToken(t *testing.T) {
	b := []byte("Pa$$w0rd")

	j := &webtoken.JWTAuth{
		Clock:                webtoken.Clock{},
		PrivateKey:           &b,
		RefreshTokenDuration: time.Hour * 24 * 7,
		AccessTokenDuration:  time.Hour * 2,
	}

	s, err := j.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsInVzZXJJRCI6Mn0.rN-m8KD3oLK_-kW3vKaialQ06Fy_k_d95cPQJyalTJE")
	assert.Equal(t, "", s)
	assert.NotEqual(t, nil, err)
}

func TestGoodGenerateVerifyTokenReal(t *testing.T) {
	clock := new(CustomClock)
	clock.Set(time.Date(2006, 01, 15, 12, 34, 0, 0, time.UTC))

	b := []byte("s3cr3tp@$$w0rdk33")

	j := &webtoken.JWTAuth{
		Clock:                clock,
		PrivateKey:           &b,
		RefreshTokenDuration: time.Hour * 24 * 7,
		AccessTokenDuration:  time.Hour * 2,
	}

	s, err := j.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjExMzczMjg0NDUsInVzZXJJRCI6MTAwfQ.ShRjs6DHypHAMf-H1a4opeZ3aEF37quOpLMnfxAXUlU")
	assert.Equal(t, "100", s)
	assert.Equal(t, nil, err)
}

func TestMissingExpiration(t *testing.T) {
	clock := new(CustomClock)
	clock.Set(time.Date(2006, 01, 15, 12, 34, 0, 0, time.UTC))

	b := []byte("s3cr3tp@$$w0rdk33")

	j := &webtoken.JWTAuth{
		Clock:                webtoken.Clock{},
		PrivateKey:           &b,
		RefreshTokenDuration: time.Hour * 24 * 7,
		AccessTokenDuration:  time.Hour * 2,
	}

	s, err := j.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsInVzZXJJRCI6Mn0.rN-m8KD3oLK_-kW3vKaialQ06Fy_k_d95cPQJyalTJE")
	assert.Equal(t, "", s)
	assert.Equal(t, webtoken.ErrMissingExpiration, err)

	r, err := j.RefreshAccessToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsInVzZXJJRCI6Mn0.rN-m8KD3oLK_-kW3vKaialQ06Fy_k_d95cPQJyalTJE")
	assert.Nil(t, r)
	assert.Equal(t, webtoken.ErrMissingExpiration, err)
}

// CustomerClock implements the IClock interface.
type CustomClock struct {
	t time.Time
}

// Now sets the current time.
func (c *CustomClock) Set(t time.Time) {
	c.t = t
}

// Now returns the current time.
func (c *CustomClock) Now() time.Time {
	return c.t
}

func TestTokenExpirationSameTime(t *testing.T) {
	clock := new(CustomClock)
	clock.Set(time.Date(2006, 01, 15, 12, 34, 0, 0, time.UTC))

	bb := []byte("Pa$$w0rd")

	j := &webtoken.JWTAuth{
		Clock:                clock,
		PrivateKey:           &bb,
		RefreshTokenDuration: time.Second * 5,
		AccessTokenDuration:  time.Second * 5,
	}

	u := new(webtoken.User)
	u.ID = "100"

	a, r, err := j.GenerateTokens(u)

	assert.NotNil(t, a)
	assert.NotNil(t, r)
	assert.Equal(t, nil, err)

	s, err := j.Verify(a.Token)
	assert.Equal(t, "100", s)
	assert.Equal(t, nil, err)
}

func TestTokenExpirationAfter(t *testing.T) {
	clock := new(CustomClock)
	clock.Set(time.Date(2006, 01, 15, 12, 34, 0, 0, time.UTC))

	bb := []byte("Pa$$w0rd")

	j := &webtoken.JWTAuth{
		Clock:                clock,
		PrivateKey:           &bb,
		RefreshTokenDuration: time.Second * 5,
		AccessTokenDuration:  time.Second * 5,
	}

	u := new(webtoken.User)
	u.ID = "100"

	a, r, err := j.GenerateTokens(u)

	assert.NotNil(t, a)
	assert.NotNil(t, r)
	assert.Equal(t, nil, err)

	// Set the clock to exact expiration a year prior.
	clock.Set(time.Date(2005, 01, 15, 12, 34, 5, 0, time.UTC))
	s, err := j.Verify(a.Token)
	assert.Equal(t, "100", s)
	assert.Equal(t, nil, err)

	// Set the clock to exact expiration.
	clock.Set(time.Date(2006, 01, 15, 12, 34, 5, 0, time.UTC))
	s, err = j.Verify(a.Token)
	assert.Equal(t, "100", s)
	assert.Equal(t, nil, err)

	// Set the clock to a second after the expiration.
	clock.Set(time.Date(2006, 01, 15, 12, 34, 6, 0, time.UTC))
	s, err = j.Verify(a.Token)
	assert.Equal(t, "", s)
	assert.NotEqual(t, nil, err)
}

func TestTokenRefresh(t *testing.T) {
	clock := new(CustomClock)
	clock.Set(time.Date(2006, 01, 15, 12, 34, 0, 0, time.UTC))

	bb := []byte("Pa$$w0rd")

	j := &webtoken.JWTAuth{
		Clock:                clock,
		PrivateKey:           &bb,
		RefreshTokenDuration: time.Second * 5,
		AccessTokenDuration:  time.Second * 5,
	}

	u := new(webtoken.User)
	u.ID = "100"

	a, r, err := j.GenerateTokens(u)

	assert.NotNil(t, a)
	assert.NotNil(t, r)
	assert.Equal(t, nil, err)

	b, err := j.RefreshAccessToken(a.Token)
	assert.NotEqual(t, nil, b)
	assert.Equal(t, nil, err)

	// Ensure the user data is still available in the new token.
	s, err := j.Verify(b.Token)
	assert.Equal(t, "100", s)
	assert.Equal(t, nil, err)
}

func TestTokenRefreshFail(t *testing.T) {
	clock := new(CustomClock)
	clock.Set(time.Date(2006, 01, 15, 12, 34, 0, 0, time.UTC))

	bb := []byte("Pa$$w0rd")

	j := &webtoken.JWTAuth{
		Clock:                clock,
		PrivateKey:           &bb,
		RefreshTokenDuration: time.Second * 5,
		AccessTokenDuration:  time.Second * 5,
	}

	u := new(webtoken.User)
	u.ID = "100"

	a, r, err := j.GenerateTokens(u)

	assert.NotNil(t, a)
	assert.NotNil(t, r)
	assert.Equal(t, nil, err)

	b, err := j.RefreshAccessToken("bad")
	assert.Nil(t, b)
	assert.NotEqual(t, nil, err)
}

func TestTokenBadSecret(t *testing.T) {
	clock := new(CustomClock)
	clock.Set(time.Date(2006, 01, 15, 12, 34, 0, 0, time.UTC))

	j := &webtoken.JWTAuth{
		Clock:                clock,
		PrivateKey:           nil,
		RefreshTokenDuration: time.Second * 5,
		AccessTokenDuration:  time.Second * 5,
	}

	u := new(webtoken.User)
	u.ID = "100"

	_, _, err := j.GenerateTokens(u)
	assert.Equal(t, webtoken.ErrBadSecret, err)

	_, err = j.RefreshAccessToken("bad")
	assert.Equal(t, webtoken.ErrBadSecret, err)

	_, err = j.Verify("bad")
	assert.Equal(t, webtoken.ErrBadSecret, err)
}
