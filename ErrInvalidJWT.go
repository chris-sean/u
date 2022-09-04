// Generated by ESG at 2022-09-04 21:18:50. github.com/simplefelix/esg

package u

import "fmt"

type InvalidJWT struct {
	_extra_ interface{}
	err interface{}
}

// ErrorCode change it as you prefer.
func (e InvalidJWT) ErrorCode() interface{} {
	return "InvalidJWT"
}

// StatusCode refers to http response status code.
// Developer may want to set response status code based on error.
// For example, if the error is caused by bad request, then change the return value to 400.
// Ignore this function if no need for your project.
func (e InvalidJWT) StatusCode() int {
	return 500
}

// Extra returns _extra_ which can be set by user. Usage of _extra_ is determined by user.
func (e InvalidJWT) Extra() interface{} {
	return e._extra_
}

// SetExtra sets _extra_ with a value by user. Usage of _extra_ is determined by user.
func (e *InvalidJWT) SetExtra(extra interface{}) {
	e._extra_ = extra
}

// Error implementation to error interface.
func (e InvalidJWT) Error() string {
	return fmt.Sprintf(`%v`, e.err)
}

// ErrInvalidJWT is convenient constructor.
func ErrInvalidJWT(err interface{}) InvalidJWT {
	return InvalidJWT{
		err: err,
	}
}
