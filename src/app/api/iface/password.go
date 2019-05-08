package iface

// IPassword provides password hashing.
type IPassword interface {
	HashString(password string) (string, error)
	MatchString(hash, password string) bool
}
