package testutil

// MockLogger is a mocked logger.
type MockLogger struct {
	FatalfFunc func(format string, v ...interface{})
	PrintfFunc func(format string, v ...interface{})
}

// NewMockLogger is a mocked logger that displays nothing by default.
func NewMockLogger() *MockLogger {
	return &MockLogger{
		FatalfFunc: func(format string, v ...interface{}) {},
		PrintfFunc: func(format string, v ...interface{}) {},
	}
}

// Fatalf .
func (l *MockLogger) Fatalf(format string, v ...interface{}) {
	l.FatalfFunc(format, v...)
}

// Printf .
func (l *MockLogger) Printf(format string, v ...interface{}) {
	l.PrintfFunc(format, v...)
}
