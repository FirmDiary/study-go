package v1

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	//NumField() return the number of fields in the value
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		//Kind() return the type of the field
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
