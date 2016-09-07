package display

import (
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	DisplayN(name, x, -1)
}

func DisplayN(name string, x interface{}, maxDepth int) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 0, maxDepth)
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Array:
		s := "["
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				s += ","
			}
			s += formatAtom(v.Index(i))
		}
		return s + "]"
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		s := fmt.Sprintf("%s{", v.Type().Name())
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				s += ","
			}
			s += fmt.Sprintf("%s = %s", v.Type().Field(i).Name, formatAtom(v.Field(i)))
		}
		return s + "}"
	default:
		return v.Type().String() + " value"
	}
}

func display(path string, v reflect.Value, depth, maxDepth int) {
	if maxDepth >= 0 && depth > maxDepth {
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), depth + 1, maxDepth)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), depth + 1, maxDepth)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key), depth + 1, maxDepth)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), depth, maxDepth)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), depth + 1, maxDepth)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
