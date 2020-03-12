package testutil

// MockToken is a mocked webtoken.
type MockToken struct {
	GenerateFunc GenerateFuncType
	VerifyValue  string
	VerifyError  error
}

// GenerateFuncType .
type GenerateFuncType func(userID string) (string, error)

// GenerateFuncDefault .
var GenerateFuncDefault = func(userID string) (string, error) {
	return "", nil
}

// Generate .
func (mt *MockToken) Generate(userID string) (string, error) {
	if mt.GenerateFunc != nil {
		return mt.GenerateFunc(userID)
	}
	return GenerateFuncDefault(userID)
}

// Verify .
func (mt *MockToken) Verify(s string) (string, error) {
	return mt.VerifyValue, mt.VerifyError
}
