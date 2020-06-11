package src

import (
	"reflect"
	"strings"
)

func getValue(st interface{}) reflect.Value {
	vIface := reflect.ValueOf(st)

	// Check if the passed interface is a pointer
	if vIface.Type().Kind() != reflect.Ptr {
		vIface = reflect.New(reflect.TypeOf(st))
	}
	return vIface
}

func HasField(st interface{}, ex string) bool {
	i := st
	elem := getValue(i).Elem()
	for _, token := range strings.Split(ex, ".") {
		field := elem.FieldByName(strings.Title(token))
		if !field.IsValid() {
			return false
		}
		kind := field.Type()
		if kind.Kind() == reflect.Ptr {
			// dereferencing a nil pointer will always return reflect.Invalid. The solution is to create a new instance first
			elem = reflect.New(field.Type().Elem()).Elem()
		}
		if kind.Kind() == reflect.Struct {
			elem = getValue(field.Interface()).Elem()
		}
	}
	return true
}
