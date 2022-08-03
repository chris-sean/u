// Generated by ESG at 2021-09-20 11:54:59. github.com/simplefelix/esg

package u

import "fmt"

type GRPCDialErr struct {
	_extra_ interface{}
	host    interface{}
	err     interface{}
}

// ErrorCode change it as you prefer.
func (e GRPCDialErr) ErrorCode() interface{} {
	return "GRPCDialErr"
}

// StatusCode refers to http response status code.
// Developer may want to set response status code based on error.
// For example, if the error is caused by bad request, then change the return value to 400.
// Ignore this function if no need for your project.
func (e GRPCDialErr) StatusCode() int {
	return 500
}

// Extra returns _extra_ which can be set by user. Usage of _extra_ is determined by user.
func (e GRPCDialErr) Extra() interface{} {
	return e._extra_
}

// SetExtra sets _extra_ with a value by user. Usage of _extra_ is determined by user.
func (e GRPCDialErr) SetExtra(extra interface{}) {
	e._extra_ = extra
}

// Error implementation to error interface.
func (e GRPCDialErr) Error() string {
	return fmt.Sprintf(`Can't dial to grpc server %v. error=%v`, e.host, e.err)
}

// ErrGRPCDialErr is convenient constructor.
func ErrGRPCDialErr(host, err interface{}) GRPCDialErr {
	return GRPCDialErr{
		host: host,
		err:  err,
	}
}