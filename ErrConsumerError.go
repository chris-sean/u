// Generated by ESG at 2021-11-05 17:13:20. github.com/simplefelix/esg

package u

import "fmt"

type ConsumerError struct {
	_extra_ interface{}
	err     interface{}
}

// ErrorCode change it as you prefer.
func (e ConsumerError) ErrorCode() interface{} {
	return "ConsumerError"
}

// StatusCode refers to http response status code.
// Developer may want to set response status code based on error.
// For example, if the error is caused by bad request, then change the return value to 400.
// Ignore this function if no need for your project.
func (e ConsumerError) StatusCode() int {
	return 500
}

// Extra returns _extra_ which can be set by user. Usage of _extra_ is determined by user.
func (e ConsumerError) Extra() interface{} {
	return e._extra_
}

// SetExtra sets _extra_ with a value by user. Usage of _extra_ is determined by user.
func (e ConsumerError) SetExtra(extra interface{}) {
	e._extra_ = extra
}

// Error implementation to error interface.
func (e ConsumerError) Error() string {
	return fmt.Sprintf(`%v`, e.err)
}

// ErrConsumerError is convenient constructor.
func ErrConsumerError(err interface{}) ConsumerError {
	return ConsumerError{
		err: err,
	}
}