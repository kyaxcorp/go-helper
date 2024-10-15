package _struct

import "github.com/kyaxcorp/go-helper/_struct/defaults"

func SetDefaultValues(obj interface{}) error {
	return defaults.Set(obj)
}
