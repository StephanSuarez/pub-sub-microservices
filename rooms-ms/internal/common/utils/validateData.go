package utils

import (
	"reflect"
)

func ValidateData(roomDto interface{}) string {
	v := reflect.ValueOf(roomDto)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if value.Kind() == reflect.String && value.String() == "" {
			return field.Name + " can not be empty"
		}
	}

	return ""
}
