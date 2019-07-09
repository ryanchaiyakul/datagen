package httplib

import (
	"fmt"
	"strings"
)

var imaginaryDenotation string

func stringComplex(v interface{}) interface{} {
	if imaginaryDenotation == "" {
		imaginaryDenotation = "i"
	}
	var ret interface{}
	switch v.(type) {
	case complex128:
		v = strings.Replace(fmt.Sprint(v), "i", imaginaryDenotation, 1)
	case []complex128:
		ret = strings.Replace(strings.Replace(fmt.Sprint(v), "i", imaginaryDenotation, -1), " ", ",", -1)
	case [][]complex128:
		ret = strings.Replace(strings.Replace(fmt.Sprint(v), "i", imaginaryDenotation, -1), " ", ",", -1)
	case [][][]complex128:
		ret = strings.Replace(strings.Replace(fmt.Sprint(v), "i", imaginaryDenotation, -1), " ", ",", -1)
	case [][][][]complex128:
		ret = strings.Replace(strings.Replace(fmt.Sprint(v), "i", imaginaryDenotation, -1), " ", ",", -1)
	default:
		ret = v
	}
	return ret
}
