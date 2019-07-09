package httplib

import (
	"github.com/ryanchaiyakul/datagen/internal/genlib"
)

// HTTPBoolParams extends ComplexParams for HTTP parameter insertion
type HTTPBoolParams struct {
	*genlib.BoolParams
}

func init() {
	HTTPParams["bool"] = &HTTPBoolParams{&genlib.BoolParams{}}
}

// SetParams allows for setting of parameters in HTTPComplexParams
func (curParams *HTTPBoolParams) SetParams(k string, v string) error {
	return nil
}

// New returns an empty object of the same config type
func (curParams *HTTPBoolParams) New() DataGenHTTP {
	return &HTTPBoolParams{&genlib.BoolParams{}}
}
