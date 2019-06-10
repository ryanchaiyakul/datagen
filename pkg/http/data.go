package httplib

import (
	"errors"
	"reflect"
)

// GenData returns a permutation of the given data type
// the data type can be user defined (eg. struct) but it must be primitive at the lowest level
// INCOMPLETE DOES NOT FUNCTION
func GenData(input interface{}) (ret interface{}, err error) {
	switch input.(type) {
	case reflect.Value:
		switch v := input.(reflect.Value); v.Kind() {
		case reflect.Ptr:
			if _, newErr := GenData(v.Elem()); newErr == nil {
				ret = input
			} else {
				err = newErr
			}
		case reflect.String:
			if v.CanSet() {
				input.(reflect.Value).SetString("hello")
			} else {
				ret = reflect.ValueOf("hello")
			}
		case reflect.Struct:
			for i := 0; i < v.NumField(); i++ {
				if field := v.FieldByIndex([]int{i}); field.IsValid() {
					if retField, newErr := GenData(field); retField != nil {
						ret, err = nil, errors.New("GenData : struct should be passed in as a pointer")
					} else if newErr != nil {
						err = newErr
					} else {
						ret = input
					}
				}
			}
		default:
			ret, err = nil, errors.New("GenData : data type generation not implemented")
		}
	default:
		retReflectValue, newErr := GenData(reflect.ValueOf(input))
		if newErr == nil {
			ret = retReflectValue.(reflect.Value).Interface()
		} else {
			err = newErr
		}
	}
	return ret, err
}
