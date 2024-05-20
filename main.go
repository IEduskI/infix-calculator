package main

import (
	"calculator-svc/conversion"
	"calculator-svc/evaluate"
	"fmt"
)

func main() {
	//Create the reader instance
	//reader := bufio.NewReader(os.Stdin)

	var infix string
	fmt.Print("Enter the infix expression: ")
	fmt.Scanln(&infix)

	postfix, _ := conversion.InfixToPostfix(infix)

	fmt.Println("Postfix expression: " + postfix.String())

	result, _ := evaluate.Postfix(postfix.String())

	fmt.Printf("Result: %.2f", result)
}
