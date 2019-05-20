package mock

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
)

// pop will return a value from the array.
func (m *Mocker) pop(i interface{}) {
	// Check for errors.
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		err := errors.New("must pass a pointer, not a value")
		log.Println(err)
		return
	}

	// Get the caller.
	caller, ok := m.findCaller(4)
	if !ok {
		err := errors.New("could not find caller")
		log.Println(err)
		return
	}

	// Set the interface.
	v.Elem().Set(reflect.ValueOf(m.items[caller][0]))

	// Remove the value.
	m.items[caller] = remove(m.items[caller], 0)
	if len(m.items[caller]) == 0 {
		delete(m.items, caller)
	}
}

// findCaller returns the caller and if there are items to process for the
// caller.
func (m *Mocker) findCaller(i int) (string, bool) {
	caller := getCaller(i)
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
