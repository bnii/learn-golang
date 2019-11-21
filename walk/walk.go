package walk

import "reflect"

func Walk(iff interface{}, fn func(s string)) {
	val := getVal(iff)

	walkValue := func(val reflect.Value) {
		Walk(val.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}
}

func getVal(iff interface{}) reflect.Value {
	val := reflect.ValueOf(iff)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
