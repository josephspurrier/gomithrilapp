package passhash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStringString tests string to string hash.
func TestStringString(t *testing.T) {
	ph := New()
	plainText := "This is a test."
	hash, err := ph.Hash(plainText)
	assert.Nil(t, err)
	assert.True(t, ph.Match(hash, plainText))

	plainText2 := "This is a test2."
	hash, err = ph.Hash(plainText2)
	assert.Nil(t, err)
	assert.False(t, ph.Match(hash, plainText))
}

// TestByteByte tests byte to byte hash.
func TestByteByte(t *testing.T) {
	ph := New()
	plainText := []byte("This is a test.")
	hash, err := ph.HashBytes(plainText)
	assert.Nil(t, err)
	assert.True(t, ph.MatchBytes(hash, plainText))

	plainText2 := []byte("This is a test2.")
	hash, err = ph.HashBytes(plainText2)
	assert.Nil(t, err)
	assert.False(t, ph.MatchBytes(hash, plainText))
}

// TestStringByte tests string to byte hash.
func TestStringByte(t *testing.T) {
	ph := New()
	plainText := "This is a test."
	hash, err := ph.Hash(plainText)
	assert.Nil(t, err)
	assert.True(t, ph.MatchBytes([]byte(hash), []byte(plainText)))
}

// TestByteString tests byte to string hash.
func TestByteString(t *testing.T) {
	ph := New()
	plainText := []byte("This is a test.")
	hash, err := ph.HashBytes(plainText)
	assert.Nil(t, err)
	assert.True(t, ph.Match(string(hash), string(plainText)))
}

// TestHashEmpty tests empty string which should pass fine.
func TestHashEmpty(t *testing.T) {
	ph := New()
	plainText := ""
	hash, err := ph.Hash(plainText)
	assert.Nil(t, err)
	assert.True(t, ph.Match(hash, plainText))
}
