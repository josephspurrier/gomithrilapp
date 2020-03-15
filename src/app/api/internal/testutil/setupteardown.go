package testutil

import (
	"app/api/endpoint"
	"app/api/pkg/database"
)

// CoreTest is a collection of utilities for testing.
type CoreTest struct {
	DB      *database.DBW
	Core    endpoint.Core
	Test    *Mocks
	Request *Request
}

// Setup will set up the test utilities.
func Setup() *CoreTest {
	ct := new(CoreTest)
	ct.DB = LoadDatabase(ct.Core.Log)
	ct.Core, ct.Test = Services(ct.DB)
	ct.Request = NewRequest()
	return ct
}

// Teardown will teardown the test utilities - should be called with a defer.
func (ct *CoreTest) Teardown() {
	TeardownDatabase(ct.DB)
}
