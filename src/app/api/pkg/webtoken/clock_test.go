package webtoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	assert.True(t, time.Since(new(clock).Now()) < time.Second)
}
