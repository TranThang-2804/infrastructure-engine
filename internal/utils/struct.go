package utils

import "reflect"

func GetStructName(i any) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}
