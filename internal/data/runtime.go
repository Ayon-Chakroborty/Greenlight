package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// if UnmarshalJSON is unable to parse string then return this error
var ErrInvalidRtuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r *Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// get rid of quotes from "<runtime> mins"
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRtuntimeFormat
	}

	// check if string is only 2 parts "<runtime> mins"
	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins"{
		return ErrInvalidRtuntimeFormat
	}

	// take the runtime value and convert it to int
	val, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil{
		return ErrInvalidRtuntimeFormat
	}

	// set receiver to the runtime value from request body
	*r = Runtime(val)

	return nil
}
