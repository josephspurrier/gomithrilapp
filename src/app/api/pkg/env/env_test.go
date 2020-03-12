package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Connection holds the details for a database connection.
type Connection struct {
	Username string `json:"Username" env:"DB_USERNAME"`
	Password string `json:"Password" env:"DB_PASSWORD"`
	Name     string `json:"Database" env:"DB_NAME"`
	Port     int    `json:"Port" env:"DB_PORT"`
	SSL      bool   `json:"SSL" env:"DB_SSL"`
}

// ConnectionDefaults holds the details for a database connection with defaults.
type ConnectionDefaults struct {
	Username string `json:"Username" env:"DB_USERNAME"`
	Password string `json:"Password" env:"DB_PASSWORD"`
	Name     string `json:"Database" env:"DB_NAME" default:"foo"`
	Port     int    `json:"Port" env:"DB_PORT" default:"8081"`
	SSL      bool   `json:"SSL" env:"DB_SSL" default:"true"`
}

func TestUnmarshalEmpty(t *testing.T) {
	c := new(Connection)
	err := Unmarshal(c, "")
	assert.Nil(t, err)

	assert.Equal(t, "", c.Username)
	assert.Equal(t, "", c.Password)
	assert.Equal(t, "", c.Name)
	assert.Equal(t, 0, c.Port)
	assert.Equal(t, false, c.SSL)
}

func TestUnmarshalFilled(t *testing.T) {
	os.Setenv("DB_USERNAME", "a")
	os.Setenv("DB_PASSWORD", "B")
	os.Setenv("DB_NAME", "c123")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_SSL", "TRUE")

	c := new(Connection)
	err := Unmarshal(c, "")
	assert.Nil(t, err)

	assert.Equal(t, "a", c.Username)
	assert.Equal(t, "B", c.Password)
	assert.Equal(t, "c123", c.Name)
	assert.Equal(t, 3306, c.Port)
	assert.Equal(t, true, c.SSL)

	os.Unsetenv("DB_USERNAME")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_SSL")
}

func TestUnmarshalDefaults(t *testing.T) {
	c := new(ConnectionDefaults)
	err := Unmarshal(c, "")
	assert.Nil(t, err)

	assert.Equal(t, "", c.Username)
	assert.Equal(t, "", c.Password)
	assert.Equal(t, "foo", c.Name)
	assert.Equal(t, 8081, c.Port)
	assert.Equal(t, true, c.SSL)
}

func TestUnmarshalFilledPrefix(t *testing.T) {
	os.Setenv("TEST_DB_USERNAME", "a")
	os.Setenv("TEST_DB_PASSWORD", "B")
	os.Setenv("TEST_DB_NAME", "c123")
	os.Setenv("TEST_DB_PORT", "3306")
	os.Setenv("TEST_DB_SSL", "TRUE")

	c := new(Connection)
	err := Unmarshal(c, "TEST_")
	assert.Nil(t, err)

	assert.Equal(t, "a", c.Username)
	assert.Equal(t, "B", c.Password)
	assert.Equal(t, "c123", c.Name)
	assert.Equal(t, 3306, c.Port)
	assert.Equal(t, true, c.SSL)

	os.Unsetenv("TEST_DB_USERNAME")
	os.Unsetenv("TEST_DB_PASSWORD")
	os.Unsetenv("TEST_DB_NAME")
	os.Unsetenv("TEST_DB_PORT")
	os.Unsetenv("TEST_DB_SSL")
}

func TestUnmarshalError(t *testing.T) {
	c := "string"
	err := Unmarshal(c, "")
	assert.Contains(t, err.Error(), "type not pointer")

	d := "string"
	err = Unmarshal(&d, "")
	assert.Contains(t, err.Error(), "type not struct")

	os.Setenv("DB_SSL", "TRUEX")
	f := new(Connection)
	err = Unmarshal(f, "")
	assert.NotNil(t, err)
	assert.Equal(t, false, f.SSL)
	os.Unsetenv("DB_SSL")

	os.Setenv("DB_PORT", "bad")
	g := new(Connection)
	err = Unmarshal(f, "")
	assert.NotNil(t, err)
	assert.Equal(t, false, g.SSL)
	os.Unsetenv("DB_PORT")
}
