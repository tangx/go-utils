package utils

import (
	"errors"
	"fmt"
	"reflect"
)

// Contains check element in slice, array or map
func Contains(elem interface{}, target interface{}) (bool, error) {

	// reflect target type
	elemType := reflect.TypeOf(elem)
	// elemValue := reflect.ValueOf(elem)
	targetType := reflect.TypeOf(target)
	targetValue := reflect.ValueOf(target)
	switch targetType.Kind() {
	case reflect.Array, reflect.Slice:
		{
			if elemType != targetType.Elem() {
				msg := fmt.Sprintf("%s does not match %s", elemType.String(), targetType.String())
				return false, errors.New(msg)
			}
			for i := 0; i < targetValue.Len(); i++ {
				if targetValue.Index(i).Interface() == elem {
					return true, nil
				}
			}
			return false, errors.New("not found")
		}

	case reflect.Map:
		{
			if elemType != targetType.Key() {
				msg := fmt.Sprintf("%s does not match %s", elemType.String(), targetType.String())
				return false, errors.New(msg)
			}

			if targetValue.MapIndex(reflect.ValueOf(elem)).IsValid() {
				return true, nil
			}
			return false, errors.New("not found")
		}
	}

	return false, errors.New("not support target")
}
