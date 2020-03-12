// Package env will fill a struct from environment variables.
package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Unmarshal will fill a struct from environment variables. It supports struct
// values of string, int, and bool.
func Unmarshal(dst interface{}, prefix string) (err error) {
	// Ensure a pointer is passed in.
	vdst := reflect.ValueOf(dst)
	if vdst.Kind() != reflect.Ptr {
		return fmt.Errorf("dst type not pointer - expected 'struct pointer' but got '%v'", vdst.Kind())
	}

	// Ensure a struct is passed in.
	vd := reflect.Indirect(reflect.ValueOf(dst))
	if vd.Kind() != reflect.Struct {
		return fmt.Errorf("dst type not struct - expected 'struct pointer' but got '%v pointer'", vd.Kind())
	}

	// Loop through each field.
	keys := vd.Type()
	for j := 0; j < vd.NumField(); j++ {
		field := keys.Field(j)
		tag := keys.Field(j).Tag

		// Get the env tag.
		envname := tag.Get("env")

		// Get the default.
		defaultValue := tag.Get("default")

		// Get the environment variable from the tag.
		envValue := os.Getenv(prefix + envname)

		// If the environment variable exists, set the value.
		val := ""
		if len(envValue) > 0 {
			val = envValue
		} else if len(defaultValue) > 0 {
			// If the default variable exists, set the value.
			val = defaultValue
		} else {
			// Skip.
			continue
		}

		vr := reflect.ValueOf(val)

		// If the types are the same, then set the field.
		if vr.Type() == field.Type {
			vd.Field(j).Set(vr)
			continue
		}

		// If the types are not the same, perform type conversion.
		f := vd.Field(j)
		switch f.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i64, err := strconv.ParseInt(val, 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(i64)
		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			f.SetBool(b)
		}
	}

	return
}
