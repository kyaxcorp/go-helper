package function

import (
	"reflect"

	"github.com/kyaxcorp/go-helper/value"
)

// IsFunc -> Checks the type
func IsFunc(v interface{}) bool {
	//if v == nil
	return reflect.TypeOf(v).Kind() == reflect.Func
}

// IsCallable -> It returns if it's a callable or not!
func IsCallable(v interface{}) bool {
	if value.IsNil(v) {
		return false
	}
	return reflect.TypeOf(v).Kind() == reflect.Func
}
