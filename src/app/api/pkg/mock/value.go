package mock

// String .
func (m *Mocker) String() (v string) {
	m.pop(&v)
	return
}

// Int .
func (m *Mocker) Int() (v int) {
	m.pop(&v)
	return
}

// Bool .
func (m *Mocker) Bool() (v bool) {
	m.pop(&v)
	return
}

// Error .
func (m *Mocker) Error() (v error) {
	m.pop(&v)
	return
}
