package errors

import (
	"encoding/json"
)

type ErrCode int

const (
	Unknown ErrCode = 0
)

// get the error code if the error is
// associated to recognizable error types
func getErrCode(err error) ErrCode {
	val, ok := err.(*Error)
	if ok {
		return ErrCode(val.Code)
	}
	return Unknown
}

// base error structure, following is the sample response
// {"code":40008,"description":"\"Bucket already exists\"","details":"Bucket already exists","retryable":false}
type Error struct {
	Code      ErrCode `json:"code,omitempty"`
	Desc      string  `json:"description,omitempty"`
	Msg       string  `json:"details,omitempty"`
	Retryable bool    `json:"retryable,omitempty"`
}

// Error() prints out the error message string
func (e Error) Error() string {
	return e.Msg
}

// Wraps the error msg
func Wrap(msg string) error {
	return &Error{
		Msg: msg,
	}
}

func ParseError(data []byte) error {
	e := &Error{}
	err := json.Unmarshal(data, e)
	if err != nil {
		return &Error{
			Msg: string(data),
		}
	}

	return e
}
