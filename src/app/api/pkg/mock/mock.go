// Package mock allows you to return specific values for functions. It will
// panic if the wrong value types are added.
package mock

import "log"

// Mocker helps test functions.
type Mocker struct {
	enable bool
	items  map[string][]interface{}
}

// New creates a Mocker.
func New(enable bool) *Mocker {
	return &Mocker{
		enable: enable,
	}
}

// SetEnable will enable the mocker.
func (m *Mocker) SetEnable(enable bool) {
	m.enable = enable
}

// ShowCaller will output the caller.
func (m *Mocker) ShowCaller() {
	log.Println(getCaller(2))
}

// Enabled returns if the mock is enabled for the function.
func (m *Mocker) Enabled() bool {
	if !m.enable || len(m.items) == 0 {
		return false
	}

	caller := getCaller(2)
	_, ok := m.items[caller]
	if !ok {
		return false
	}

	return true
}

// Add will append a value.
func (m *Mocker) Add(caller string, arr ...interface{}) {
	if m.items == nil {
		m.items = make(map[string][]interface{})
	}

	_, ok := m.items[caller]
	if !ok {
		m.items[caller] = make([]interface{}, 0)
	}

	for _, i := range arr {
		m.items[caller] = append(m.items[caller], i)
	}
}
