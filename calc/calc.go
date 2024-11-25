package calc

type Calculator interface {
	Calculate(a, b int) int
}

type Addition struct {
}

func (t *Addition) Calculate(a, b int) int {
	return a + b
}
