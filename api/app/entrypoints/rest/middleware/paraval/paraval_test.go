package paraval

import (
	"fmt"
	"testing"
)

func NewParam(key, value string) Param {
	return Param{
		Key:   key,
		Value: value,
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in     []Param
		outVal Validated
		outErr error
	}{
		{[]Param{}, Validated{}, nil},
		{[]Param{NewParam("a", "")}, Validated{}, nil},
		{[]Param{NewParam("id", "")}, Validated{}, nil},
		{[]Param{NewParam("aaidaa", "")}, Validated{}, nil},
		{[]Param{NewParam("ID", "")}, Validated{}, ErrMissingParam},
		{[]Param{NewParam("aaIDaa", "")}, Validated{}, ErrMissingParam},
		{[]Param{NewParam("ID", "a")}, Validated{}, ErrInvalidParam},
		{[]Param{NewParam("ID", "0")}, Validated{}, ErrInvalidParam},
		{
			[]Param{NewParam("ID", "10")},
			Validated{"ID": uint(10)},
			nil,
		},
		{
			[]Param{NewParam("aID", "1"), NewParam("bID", "2")},
			Validated{"aID": uint(1), "bID": uint(2)},
			nil,
		},
		{
			[]Param{NewParam("aID", "1"), NewParam("bID", "a")},
			Validated{},
			ErrInvalidParam,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%+v", c.in), func(t *testing.T) {
			t.Parallel()

			val, err := Validate(c.in)
			if fmt.Sprintf("%+v", val) != fmt.Sprintf("%+v", c.outVal) {
				t.Errorf("got %+v, want %+v", val, c.outVal)
			}
			if err != c.outErr {
				t.Errorf("got %+v, want %+v", err, c.outErr)
			}
		})
	}

}
