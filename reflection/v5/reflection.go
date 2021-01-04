package v1

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	//NumField() return the number of fields in the value
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		//Kind() return the type of the field
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		//we can't use NumField on Pointer Value
		//use Elem() extract the underlying value
		val = val.Elem()
	}

	return val
}
