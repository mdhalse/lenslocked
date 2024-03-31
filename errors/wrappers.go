package errors

import "errors"

// These variables are used to give us access to existing
// functions in the std library errors package.  We can
// also wrap them in custom functionality if needed or
// mock them during testing.
var (
	As = errors.As
	Is = errors.Is
)
