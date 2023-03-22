package goerror

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_Error(t *testing.T) {
	fmt.Println("------------------")
	err := Error("new error")
	fmt.Println(err.Error())
	fmt.Println(err.GetStackTarce())
}

func TestError_Err(t *testing.T) {
	fmt.Println("------------------")
	err := New("test", "ERR-001")

	fmt.Println(err.GetStackTrace())

	err2 := errors.New("from errors")

	err3 := FromError(err2)

	fmt.Println(err3.GetStackTrace())

	fmt.Printf("err message: %s, code: %s\n", err.Error(), err.Code())
	fmt.Printf("err3 message: %s, code: %s\n", err3.Error(), err3.Code())
}

func TestDeferError(t *testing.T) {
	fmt.Println("------------------")

	err := Error("test defer error")

	// Definir una función para manejar los errores y llamarla con "defer"
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			t.Logf("%+v", err)
		}
	}()

	// Llamar a la función que maneja los errores antes de provocar el error
	DeferError(func(err error, stackTrace string) {
		log := fmt.Sprintf("error: %s stack: %s", err.Error(), stackTrace)
		fmt.Println(log)
		t.Log(log)
	})

	// Provocar un error intencionalmente
	panic(err)
}
