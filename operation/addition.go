package operation

const (
	additionPrecedence = 1
)

type addition struct {
	OperatorBase
}

func NewAddition() addition {
	return addition{
		OperatorBase{
			precedence: additionPrecedence,
			symbol:     "+",
		},
	}
}

func (ad addition) Evaluate(a, b float64) float64 {
	return a + b
}
