package httpmainlib

import (
	"fmt"
	"strconv"
	"strings"

	genlib "github.com/ryanchaiyakul/datagen/internal/genlib"
)

// HTTPIntParams extends HTTPIntParams for HTTP parameter insertion
type HTTPIntParams struct {
	*genlib.IntParams
}

func init() {
	HTTPParams["int"] = &HTTPIntParams{&genlib.IntParams{}}
}

// SetParams allows for setting of parameters in HTTPIntParams
func (curParams *HTTPIntParams) SetParams(k string, v string) error {
	switch k {
	case "valid_values":
		for _, v := range strings.Split(v, ",") {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.ValidValues = append(curParams.ValidValues, intV)
			}
		}
	default:
		return fmt.Errorf("getIntParam : unknown parameter : %v", k)
	}
	return nil
}

// New returns an empty object of the same config type
func (curParams *HTTPIntParams) New() DataGenHTTP {
	return &HTTPIntParams{&genlib.IntParams{}}
}
