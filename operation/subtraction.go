package operation

const (
	subtractionPrecedence = 1
)

type subtraction struct {
	OperatorBase
}

func NewSubtraction() subtraction {
	return subtraction{
		OperatorBase{
			precedence: subtractionPrecedence,
			symbol:     "-",
		},
	}
}

func (s subtraction) Evaluate(a, b float64) float64 {
	return a - b
}
