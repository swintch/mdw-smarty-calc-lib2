package handler

import (
	"bytes"
	"errors"
	"strconv"
	"testing"

	"github.com/swintch/mdw-smarty-calc-lib2/calc"
)

func assertEquals(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual = %s, expected = %s", actual, expected)
	}
}

func assertError(t *testing.T, error, expected error) {
	t.Helper()
	if !errors.Is(error, expected) {
		t.Errorf("Expected error %v, but got error %v", expected, error)
	}
}

func TestHandler_NotEnoughArgs(t1 *testing.T) {
	handlerObj := NewHandler(nil, nil)
	err := handlerObj.Handle(nil)
	assertError(t1, err, requiresTwoArgs)
}

func TestHandler_InvalidFirstArgument(t1 *testing.T) {
	testValues := []string{"NaN", "3"}
	handlerObj := NewHandler(nil, nil)
	err := handlerObj.Handle(testValues)
	assertError(t1, err, InvalidArg)
	assertError(t1, err, strconv.ErrSyntax)

}

func TestHandler_InvalidSecondArgument(t1 *testing.T) {
	testValues := []string{"1", "NaN"}
	handlerObj := NewHandler(nil, nil)
	err := handlerObj.Handle(testValues)
	assertError(t1, err, InvalidArg)
	assertError(t1, err, strconv.ErrSyntax)
}

func TestHandler_ValidArguments(t1 *testing.T) {
	var output bytes.Buffer
	testValues := []string{"1", "3"}
	calculator := &calc.Addition{}
	handlerObj := NewHandler(calculator, &output)
	err := handlerObj.Handle(testValues)
	assertError(t1, err, nil)
	assertEquals(t1, output.String(), "4")
}

func TestHandler_WriterError(t1 *testing.T) {
	boink := errors.New("boink")
	output := WriterErr{err: boink}
	testValues := []string{"1", "3"}
	calculator := &calc.Addition{}
	handlerObj := NewHandler(calculator, &output)
	err := handlerObj.Handle(testValues)
	assertError(t1, err, boink)
	assertError(t1, err, WriterError)
}

type WriterErr struct {
	err error
}

func (this *WriterErr) Write(p []byte) (n int, err error) {
	return 0, this.err
}
