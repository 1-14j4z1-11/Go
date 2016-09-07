package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"regexp"
)

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]reflect.Value)
	formats := make(map[string]string)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
		if format := tag.Get("format"); format != "" {
			formats[name] = format
			fmt.Printf("Set : %s : %s", name, format)
		}
	}

	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value, formats[name]); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value, formats[name]); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string, format string) error {
	switch v.Kind() {
	case reflect.String:
		reg, err := regexp.Compile("^" + format + "$")
		if err != nil {
			return fmt.Errorf("Invalid 'format' field tag : %s", format)
		}
		if !reg.MatchString(value) {
			return fmt.Errorf("Invalid format field value : %s (Required format : %s)", value, format)
		}
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
