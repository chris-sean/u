// Generated by ESG at 2021-09-29 09:54:09. github.com/simplefelix/esg

package u

import "fmt"

type ShortUUIDLenConstraint struct {
	_extra_ interface{}
}

// ErrorCode change it as you prefer.
func (e ShortUUIDLenConstraint) ErrorCode() interface{} {
	return "ShortUUIDLenConstraint"
}

// StatusCode refers to http response status code.
// Developer may want to set response status code based on error.
// For example, if the error is caused by bad request, then change the return value to 400.
// Ignore this function if no need for your project.
func (e ShortUUIDLenConstraint) StatusCode() int {
	return 500
}

// Extra returns _extra_ which can be set by user. Usage of _extra_ is determined by user.
func (e ShortUUIDLenConstraint) Extra() interface{} {
	return e._extra_
}

// SetExtra sets _extra_ with a value by user. Usage of _extra_ is determined by user.
func (e ShortUUIDLenConstraint) SetExtra(extra interface{}) {
	e._extra_ = extra
}

// Error implementation to error interface.
func (e ShortUUIDLenConstraint) Error() string {
	return fmt.Sprintf(`length must be in [1, 32]`)
}

// ErrShortUUIDLenConstraint is convenient constructor.
func ErrShortUUIDLenConstraint() ShortUUIDLenConstraint {
	return ShortUUIDLenConstraint{}
}
