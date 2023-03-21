package main

import (
	"fmt"
	"strconv"
)

// Operator represents an operator, its precedence, and its associativity.
type Operator struct {
	Character  string
	Precedence int
	LeftAssoc  bool
	Eval       func(a, b float64) float64
}

// Operators map.
var operators = map[string]Operator{
	// Define the supported operators, their precedence, associativity, and evaluation functions.
	"+": {"+", 1, true, func(a, b float64) float64 { return a + b }},
	"-": {"-", 1, true, func(a, b float64) float64 { return a - b }},
	"*": {"*", 2, true, func(a, b float64) float64 { return a * b }},
	"/": {"/", 2, true, func(a, b float64) float64 { return a / b }},
}

// higherPrecedence checks the precedence and associativity for operators.
func higherPrecedence(op1, op2 string) bool {
	o1 := operators[op1]
	o2 := operators[op2]
	return (o1.LeftAssoc && o1.Precedence >= o2.Precedence) || (!o1.LeftAssoc && o1.Precedence > o2.Precedence)
}

// isOperator checks if a given token is an operator.
func isOperator(token string) bool {
	_, ok := operators[token]
	return ok
}

// tokenize takes an input string and returns a slice of tokens.
func tokenize(input string) []string {
	var tokens []string
	var buffer string

	// Iterate over each character in the input string.
	for _, char := range input {
		token := string(char)

		if isOperator(token) || token == "(" || token == ")" {
			if buffer != "" {
				tokens = append(tokens, buffer)
				buffer = ""
			}
			tokens = append(tokens, token)
		} else if token != " " {
			buffer += token
		}
	}

	if buffer != "" {
		tokens = append(tokens, buffer)
	}

	return tokens
}

// ShuntingYard implements the Shunting Yard Algorithm.
func ShuntingYard(input string) (output []string, err error) {
	var stack []string

	// Tokenize the input string.
	tokens := tokenize(input)

	// Iterate through the tokens.
	for _, token := range tokens {
		// If the token is an operator, handle operator precedence and associativity.
		if op, ok := operators[token]; ok {
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if _, ok := operators[top]; ok && higherPrecedence(top, token) {
					output = append(output, top)
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			stack = append(stack, op.Character)
		} else if token == "(" {
			// If the token is an open parenthesis, push it onto the stack.
			stack = append(stack, token)
		} else if token == ")" {
			// If the token is a close parenthesis, pop operators from the stack and add them to the output
			// until an open parenthesis is encountered.
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] != "(" {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1]
		} else {
			// If the token is a number, add it to the output.
			output = append(output, token)
		}
	}

	// Pop any remaining operators from the stack and add them to the output.
	for len(stack) > 0 {
		op := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if op == "(" || op == ")" {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		output = append(output, op)
	}

	return output, nil
}

// EvaluatePostfix evaluates a postfix expression and returns the result.
func EvaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	// Iterate through the postfix expression.
	for _, token := range postfix {
		// If the token is an operator, pop two values from the stack, apply the operator, and push the result back onto the stack.
		if op, ok := operators[token]; ok {
			if len(stack) < 2 {
				return 0, fmt.Errorf("insufficient values for operator '%s'", token)
			}

			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result := op.Eval(a, b)
			stack = append(stack, result)
		} else {
			// If the token is a number, push it onto the stack.
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", token)
			}
			stack = append(stack, num)
		}
	}

	// The final result should be the only value remaining on the stack.
	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid postfix expression")
	}

	return stack[0], nil
}

// Calculate takes an input string, converts it to postfix, and evaluates it.
func Calculate(input string) (float64, error) {
	postfix, err := ShuntingYard(input)
	if err != nil {
		return 0, err
	}

	result, err := EvaluatePostfix(postfix)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	expressions := []string{
		"3+5",
		"7-3+4",
		"3*(3+11-4)/2",
	}

	for _, exp := range expressions {
		result, err := Calculate(exp)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		fmt.Printf("Input: \"%s\"\nOutput: %.0f\n\n", exp, result)
	}
}
