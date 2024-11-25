package handler

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/swintch/mdw-smarty-calc-lib2/calc"
)

type Handler struct {
	calculator calc.Calculator
	stdout     io.Writer
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return requiresTwoArgs
	}
	integerValue1, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %w", InvalidArg, err)
	}
	integerValue2, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: %w", InvalidArg, err)
	}
	result := this.calculator.Calculate(integerValue1, integerValue2)
	_, err = fmt.Fprintf(this.stdout, "%d", result)
	if err != nil {
		return fmt.Errorf("%w: %w", WriterError, err)
	}
	return nil

}

func NewHandler(calculator calc.Calculator, stdout io.Writer) *Handler {
	return &Handler{
		calculator: calculator,
		stdout:     stdout,
	}
}

var (
	requiresTwoArgs = errors.New("Requires two arguments!")
	InvalidArg      = errors.New("Invalid argument!")
	WriterError     = errors.New("Writer error!")
)
