package structure

import (
	"reflect"
	"unsafe"
)

type checker struct {
	ptr unsafe.Pointer
	tp  reflect.Type
}

func IsCircular(x interface{}) bool {
	return circular(reflect.ValueOf(&x).Elem(), make(map[checker]bool))
}

func circular(x reflect.Value, seen map[checker]bool) bool {
	if x.CanAddr() {
		ptr := unsafe.Pointer(x.UnsafeAddr())
		c := checker{ptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}
	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return circular(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if circular(x.Index(i), seen) {
				return true
			}
		}

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if circular(x.Field(i), seen) {
				return true
			}
		}

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if circular(x.MapIndex(k), seen) {
				return true
			}
		}
	}
	return false
}
