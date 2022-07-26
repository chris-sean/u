// Package errors
// Generated by ESG at 2021-09-09 11:52:38. github.com/simplefelix/esg
package u

import "fmt"

type ParamBindingErr struct {
	_extra_ interface{}
	err     interface{}
}

func (e ParamBindingErr) ErrorCode() interface{} {
	return "ParamBindErr"
}

func (e ParamBindingErr) StatusCode() int {
	return 400
}

// Extra returns _extra_ which can be set by user. Usage of _extra_ is determined by user.
func (e ParamBindingErr) Extra() interface{} {
	return e._extra_
}

// SetExtra sets _extra_ with a value by user. Usage of _extra_ is determined by user.
func (e ParamBindingErr) SetExtra(extra interface{}) {
	e._extra_ = extra
}

func (e ParamBindingErr) Error() string {
	return fmt.Sprintf("%v", e.err)
}

func ErrParamBindingErr(err interface{}) ParamBindingErr {
	return ParamBindingErr{
		err: err,
	}
}
