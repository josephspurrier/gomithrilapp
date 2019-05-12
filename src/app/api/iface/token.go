package iface

import "time"

// IToken provides outputs for the JWT.
type IToken interface {
	Generate(userID string, duration time.Duration) (string, error)
}
