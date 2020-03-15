package testutil

// MockToken is a mocked webtoken.
type MockToken struct {
	GenerateFunc func(userID string) (string, error)
	VerifyValue  string
	VerifyError  error
}

// NewMockToken returns a new mock token.
func NewMockToken() *MockToken {
	return &MockToken{
		GenerateFunc: func(userID string) (string, error) {
			return "", nil
		},
	}
}

// Generate .
func (mt *MockToken) Generate(userID string) (string, error) {
	return mt.GenerateFunc(userID)
}

// Verify .
func (mt *MockToken) Verify(s string) (string, error) {
	return mt.VerifyValue, mt.VerifyError
}
