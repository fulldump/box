package box

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, obtained, expected interface{}) bool {
	if reflect.DeepEqual(expected, obtained) {
		return true
	}

	t.Errorf("\nExpected: %#v\nObtained: %#v\n\n", expected, obtained)
	return false
}
