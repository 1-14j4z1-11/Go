package json

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		// 出力しない

	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "\"%f+%fi\"", real(v.Complex()), imag(v.Complex()))

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)

	case reflect.Slice, reflect.Array:
		buf.WriteByte('[')
		indent := getIndent(buf)
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(",\n")
				outputIndent(buf, indent)
			}
			if err := encode(buf, v.Index(i), indent+1); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct:
		buf.WriteByte('{')
		indent := getIndent(buf)
		isFirst := true
		for i := 0; i < v.NumField(); i++ {
			if isNil(v.Field(i)) {
				continue
			}
			if !isFirst {
				buf.WriteString(",\n")
				outputIndent(buf, indent)
			}
			fmt.Fprintf(buf, "\"%s\" : ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent+1); err != nil {
				return err
			}
			isFirst = false
		}
		buf.WriteByte('}')

	case reflect.Interface:
		encode(buf, v.Elem(), indent)

	case reflect.Map:
		buf.WriteByte('{')
		indent := getIndent(buf)
		isFirst := true
		for _, key := range v.MapKeys() {
			if isNil(v.MapIndex(key)) {
				continue
			}
			if !isFirst {
				buf.WriteString(",\n")
				outputIndent(buf, indent)
			}
			if err := encode(buf, key, indent+1); err != nil {
				return err
			}
			buf.WriteString(" : ")
			if err := encode(buf, v.MapIndex(key), indent+1); err != nil {
				return err
			}
			isFirst = false
		}
		buf.WriteByte('}')

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

func getIndent(buf *bytes.Buffer) int {
	n := 0
	runes := []rune(string(buf.Bytes()))
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == '\n' {
			break
		}
		n++
	}
	return n
}

func outputIndent(buf *bytes.Buffer, indent int) {
	for i := 0; i < indent; i++ {
		buf.WriteByte(' ')
	}
}
