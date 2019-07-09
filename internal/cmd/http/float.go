package httplib

import (
	"fmt"
	"strconv"
	"strings"

	genlib "github.com/ryanchaiyakul/datagen/internal/genlib"
)

// HTTPFloatParams extends HTTPIntParams for HTTP parameter insertion
type HTTPFloatParams struct {
	*genlib.FloatParams
}

func init() {
	HTTPParams["float"] = &HTTPFloatParams{&genlib.FloatParams{}}
}

// SetParams allows for setting of parameters in HTTPIntParams
func (curParams *HTTPFloatParams) SetParams(k string, v string) error {
	switch k {
	case "valid_values":
		for _, v := range strings.Split(v, ",") {
			if floatV, err := strconv.ParseFloat(v, 64); err == nil {
				curParams.ValidValues = append(curParams.ValidValues, floatV)
			}
		}
	default:
		return fmt.Errorf("getIntParam : unknown parameter : %v", k)
	}
	return nil
}

// New returns an empty object of the same config type
func (curParams *HTTPFloatParams) New() DataGenHTTP {
	return &HTTPFloatParams{&genlib.FloatParams{}}
}
