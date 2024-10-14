package slice

import (
	"reflect"

	"github.com/kyaxcorp/go-helper/errors2/define"
)

func IndexExists(slice interface{}, indexNr int) (bool, error) {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		// panic("Invalid data-type")
		return false, define.Err(0, "invalid data-type")
	}

	return s.Len() > indexNr, nil
}
