package calc

type Calculator interface {
	Calculate(a, b int) int
}

type Addition struct{}
type Subtraction struct{}
type Multiplication struct{}
type Division struct{}

func (t *Addition) Calculate(a, b int) int {
	return a + b
}
func (t *Subtraction) Calculate(a, b int) int    { return a - b }
func (t *Multiplication) Calculate(a, b int) int { return a * b }
func (t *Division) Calculate(a, b int) int       { return a / b }
