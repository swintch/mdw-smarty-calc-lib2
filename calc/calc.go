package calc

type Addition struct {
}

func (t *Addition) Calculate(a, b int) int {
	return a + b
}
