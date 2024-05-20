package operation

const (
	multiplyPrecedence = 2
)

type multiplication struct {
	OperatorBase
}

func NewMultiplication() multiplication {
	return multiplication{
		OperatorBase{
			precedence: multiplyPrecedence,
			symbol:     "*",
		},
	}
}

func (m multiplication) Evaluate(a, b float64) float64 {
	return a * b
}
