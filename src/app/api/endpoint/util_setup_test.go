package endpoint_test

import (
	"app/api/endpoint"
	"app/api/internal/testrequest"
	"app/api/internal/testutil"
	"app/api/pkg/database"
)

// CoreTest is a collection of utilities for testing.
type CoreTest struct {
	DB      *database.DBW
	Core    endpoint.Core
	Test    *testutil.CoreTest
	Request *testrequest.TR
}

// Setup will set up the test utilities.
func Setup() *CoreTest {
	ct := new(CoreTest)
	ct.DB = testutil.LoadDatabase()
	ct.Core, ct.Test = testutil.Services(ct.DB)
	ct.Request = testrequest.New()
	return ct
}

// Teardown will teardown the test utilities - should be called with a defer.
func (ct *CoreTest) Teardown() {
	testutil.TeardownDatabase(ct.DB)
}
