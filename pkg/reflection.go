package internal

import "reflect"

func TypeName(value interface{}) string {
	if value == nil {
		return ""
	}

	t := reflect.TypeOf(value)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Name()
}
