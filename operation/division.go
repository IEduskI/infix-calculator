package operation

const (
	divisionPrecedence = 2
)

type division struct {
	OperatorBase
}

func NewDivision() division {
	return division{
		OperatorBase{
			precedence: divisionPrecedence,
			symbol:     "/",
		},
	}
}

func (d division) Evaluate(a, b float64) float64 {
	return a / b
}
