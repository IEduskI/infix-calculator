package operation

type Operator interface {
	IsGreaterOrEqual(operator Operator) bool
	Evaluate(a, b float64) float64
	Precedence() int
	Symbol() string
}

type OperatorBase struct {
	precedence int
	symbol     string
}

func (o OperatorBase) IsGreaterOrEqual(operator Operator) bool {
	return o.precedence >= operator.Precedence()
}

func (o OperatorBase) Precedence() int {
	return o.precedence
}

func (o OperatorBase) Symbol() string {
	return o.symbol
}
