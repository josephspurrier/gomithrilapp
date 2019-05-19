package mock

// String returns a string.
func (m *Mocker) String() string {
	// FIXME: Move this into the pop function.
	caller, ok := m.findCaller()
	if !ok {
		return ""
	}

	// FIXME: Change this to accept a pointer and return an error.
	return m.pop(caller).(string)
}

// Error returns an error.
func (m *Mocker) Error() error {
	// FIXME: Move this into the pop function.
	caller, ok := m.findCaller()
	if !ok {
		return nil
	}

	// FIXME: Change this to accept a pointer and return an error.
	return m.pop(caller).(error)
}
