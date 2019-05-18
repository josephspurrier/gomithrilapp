package testutil

// MockPasshash is a mocked passhash.
type MockPasshash struct {
	HashString string
	HashError  error
	MatchBool  bool
}

// Hash .
func (m *MockPasshash) Hash(password string) (string, error) {
	return m.HashString, m.HashError
}

// Match .
func (m *MockPasshash) Match(hash, password string) bool {
	return m.MatchBool
}
