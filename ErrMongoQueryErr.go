// Generated by ESG at 2021-10-15 18:05:55. github.com/simplefelix/esg

package u

import "fmt"

type MongoQueryErr struct {
	_extra_ interface{}
	err     interface{}
}

// ErrorCode change it as you prefer.
func (e MongoQueryErr) ErrorCode() interface{} {
	return "MongoQueryErr"
}

// StatusCode refers to http response status code.
// Developer may want to set response status code based on error.
// For example, if the error is caused by bad request, then change the return value to 400.
// Ignore this function if no need for your project.
func (e MongoQueryErr) StatusCode() int {
	return 500
}

// Extra returns _extra_ which can be set by user. Usage of _extra_ is determined by user.
func (e MongoQueryErr) Extra() interface{} {
	return e._extra_
}

// SetExtra sets _extra_ with a value by user. Usage of _extra_ is determined by user.
func (e *MongoQueryErr) SetExtra(extra interface{}) {
	e._extra_ = extra
}

// Error implementation to error interface.
func (e MongoQueryErr) Error() string {
	return fmt.Sprintf(`%v`, e.err)
}

// ErrMongoQueryErr is convenient constructor.
func ErrMongoQueryErr(err interface{}) MongoQueryErr {
	return MongoQueryErr{
		err: err,
	}
}