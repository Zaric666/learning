package v1

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	fieldVal := val.Field(0)
	vt := reflect.TypeOf(x)
	fieldKey := vt.Field(0)

	fn(fieldKey.Name + ":" + fieldVal.String())
}
