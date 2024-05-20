package evaluate

import (
	"calculator-svc/container"
	"calculator-svc/factory"
	"strconv"
	"strings"
)

func Postfix(postFix string) (float64, error) {
	stack := container.Stack[float64]{}
	tokens := strings.Split(postFix, " ")
	opFactory := factory.OperatorFactory{}

	for _, token := range tokens {
		operand, err := strconv.ParseFloat(token, 64)
		if err != nil {
			operator, _ := opFactory.Factory(token)

			// Last
			b, _ := stack.Pop()
			//Second last
			a, _ := stack.Pop()

			result := operator.Evaluate(a, b)

			stack.Push(result)

			continue
		}

		stack.Push(operand)
	}

	resultEvaluate, _ := stack.Pop()
	return resultEvaluate, nil
}
