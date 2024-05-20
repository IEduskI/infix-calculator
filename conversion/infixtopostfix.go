package conversion

import (
	"calculator-svc/container"
	"calculator-svc/factory"
	"calculator-svc/operation"
	"fmt"
	"strconv"
	"strings"
)

type PostFix struct {
	expression []string
}

func (p *PostFix) String() string {
	return strings.Join(p.expression, " ")
}

func (p *PostFix) Add(s string) {
	p.expression = append(p.expression, s)
}

func InfixToPostfix(s string) (PostFix, error) {
	//temArray := []string{"1", "2.5", "3", "/", "4", "*", "+"}
	//postfix := PostFix{expression: temArray}
	opFactory := factory.OperatorFactory{}
	var output []string

	tokens, _ := Tokenization(s)
	stack := container.Stack[operation.Operator]{}

	for _, token := range tokens {
		_, err := strconv.ParseFloat(token, 64)
		if err != nil {
			op, _ := opFactory.Factory(token)
			if stack.IsEmpty() {
				stack.Push(op)
				continue
			}

			length := stack.Length()

			// For to evaluate all operators
			for i := 0; i < length; i++ {
				lastOP, _ := stack.Peek()
				if op.IsGreaterOrEqual(lastOP) {
					if op.Precedence() == lastOP.Precedence() {
						lastOP, _ = stack.Pop()

						output = append(output, lastOP.Symbol())

						stack.Push(op)
						break
					}

					stack.Push(op)
					break
				}

				lastOP, _ = stack.Pop()

				output = append(output, lastOP.Symbol())
				continue
			}

			if stack.IsEmpty() {
				stack.Push(op)
				continue
			}
			continue
		}
		output = append(output, token)
	}

	if !stack.IsEmpty() {
		length := stack.Length()
		for i := 0; i < length; i++ {
			op, _ := stack.Pop()
			output = append(output, op.Symbol())
		}
	}

	postfix := PostFix{expression: output}

	fmt.Println(postfix.String())

	return postfix, nil
}

func Tokenization(s string) ([]string, error) {

	expression := strings.ReplaceAll(s, " ", "")

	var digit string
	var infixTokenized []string
	for _, v := range expression {
		if v >= '0' && v <= '9' || v == '.' {
			digit += string(v)
		} else {
			if digit != "" {
				infixTokenized = append(infixTokenized, digit)
				digit = ""
			}
			infixTokenized = append(infixTokenized, string(v))
		}
	}

	if digit != "" {
		infixTokenized = append(infixTokenized, digit)
	}

	return infixTokenized, nil
}
