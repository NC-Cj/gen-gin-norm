package core

import "fmt"

type Exception interface {
	error
}

type NoPermissions struct {
	Message string
}

type DatabaseFailure struct {
	Message string
}

type UnsupportedDataTypeError struct {
	Message string
}

// 实现 Error 接口，返回错误信息
func (e *NoPermissions) Error() string {
	return fmt.Sprintf("NoPermissions: %s", e.Message)
}

func (e *DatabaseFailure) Error() string {
	return fmt.Sprintf("DatabaseFailure: %s", e.Message)
}

func (e *UnsupportedDataTypeError) Error() string {
	return fmt.Sprintf("UnsupportedDataTypeError: %s", e.Message)
}

var ExceptionsToCatch = []Exception{
	&NoPermissions{},
	&DatabaseFailure{},
}
