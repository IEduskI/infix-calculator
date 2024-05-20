package container

import (
	"calculator-svc/operation"
	"testing"
)

func TestStack_Peek(t *testing.T) {
	stack := Stack[operation.Operator]{}

	addition := operation.NewAddition()
	stack.Push(addition)

	op, err := stack.Peek()
	if err == ErrEmptyStack {
		t.Errorf("Got: %s, but want: nil", err)
	}

	if op.Symbol() != "+" {
		t.Errorf("Got: %s, but want: +", op.Symbol())
	}

	op2, err := stack.Pop()
	if err == ErrEmptyStack {
		t.Errorf("Got: %s, but want: nil", err)
	}

	if op2.Symbol() != "+" {
		t.Errorf("Got: %s, but want: +", op2.Symbol())
	}

	_, err = stack.Peek()
	if err != ErrEmptyStack {
		t.Errorf("Got: %s, but want: nil", err)
	}

}
