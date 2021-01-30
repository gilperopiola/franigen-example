package paraval

import (
	"errors"
	"strconv"
	"strings"
)

type Param struct {
	Key   string
	Value string
}

type Validated map[string]uint

var ErrInvalidParam = errors.New("invalid param")
var ErrMissingParam = errors.New("missing param")

func Validate(params []Param) (Validated, error) {

	toValidate := make([]Param, 0, len(params))
	for _, p := range params {
		if shouldValidate(p) {
			toValidate = append(toValidate, p)
		}
	}

	validated := make(Validated)
	for _, p := range toValidate {
		if p.Value == "" {
			return nil, ErrMissingParam
		}
		v, err := strconv.ParseUint(p.Value, 10, 32)
		if err != nil {
			return nil, ErrInvalidParam
		}
		if v == 0 {
			return nil, ErrInvalidParam
		}
		validated[p.Key] = uint(v)
	}

	return validated, nil
}

func shouldValidate(param Param) bool {
	return strings.Contains(param.Key, "ID")
}
