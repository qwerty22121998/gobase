package query

import "reflect"

func isZero(value any) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Type().Kind() {
	case reflect.Slice, reflect.Chan, reflect.Array, reflect.Map:
		return v.Len() == 0
	default:
		return reflect.DeepEqual(value, reflect.Zero(v.Type()).Interface())
	}
}
