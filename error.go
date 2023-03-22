package goerror

import (
	"fmt"
	"runtime"
)

type (
	// Error is a string that implements the error interface.
	Error string

	// Err is a struct that implements the error interface.
	Err struct {
		message string
		code    string
	}
)

// Error is a string that implements the error interface.
func (e Error) Error() string {
	return string(e)
}

// GetStackTarce GetStackTrace returns the stack trace of the error.
func (e *Error) GetStackTarce() string {
	return getStackTrace()
}

// New returns a new error.
func New(message string, code ...string) (err Err) {
	err.message = message
	if len(code) > 0 {
		err.code = code[0]
	}
	return
}

// FromError returns a new error from an error.
func FromError(e error) (err Err) {
	err.message = fmt.Sprintf("%v", e)
	return
}

// ToError returns the error as an error.
func (e *Err) ToError() error {
	return fmt.Errorf("%v", e.message)
}

// Error returns the message of the error.
func (e *Err) Error() string {
	return e.message
}

// Code returns the code of the error.
func (e *Err) Code() string {
	return e.code
}

// GetStackTrace returns the stack trace of the error.
func (e *Err) GetStackTrace() string {
	return getStackTrace()
}

func getStackTrace() string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	log := ""
	for {
		frame, more := frames.Next()
		log += fmt.Sprintf("%s:%d\n", frame.File, frame.Line)
		if !more {
			break
		}
	}
	return log
}

// DeferError is a function that recovers from a panic and calls the callback function.
func DeferError(callback func(err error, stackTrace string)) {
	defer func() {
		if err := recover(); err != nil {
			recoveredErr := fmt.Errorf("%v", err)
			callback(recoveredErr, getStackTrace())
		}
	}()
}
