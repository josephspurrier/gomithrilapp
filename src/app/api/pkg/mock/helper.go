package mock

import (
	"fmt"
	"runtime"
	"strings"
)

// pop will return a value from the array.
func (m *Mocker) pop(caller string) interface{} {
	v := m.items[caller][0]
	m.items[caller] = remove(m.items[caller], 0)
	if len(m.items[caller]) == 0 {
		delete(m.items, caller)
	}
	return v
}

// findCaller returns the caller and if there are items to process for the
// caller.
func (m *Mocker) findCaller() (string, bool) {
	caller := getCaller(3)
	arr, ok := m.items[caller]
	if !ok {
		return caller, false
	}
	if len(arr) == 0 {
		return caller, false
	}

	return caller, true
}

func getCaller(i int) string {
	pc, _, _, ok := runtime.Caller(i)
	if !ok {
		return ""
	}

	f := runtime.FuncForPC(pc)
	if f == nil {
		return ""
	}

	// Turn: app/api/store.(*User).Create
	// Into this: User.Create
	arr := strings.Split(f.Name(), ".")
	file := strings.Replace(arr[1], "(", "", -1)
	file = strings.Replace(file, ")", "", -1)
	file = strings.Replace(file, "*", "", -1)

	return fmt.Sprintf("%v.%v", file, arr[2])
}

func remove(slice []interface{}, i int) []interface{} {
	return append(slice[:i], slice[i+1:]...)
}
