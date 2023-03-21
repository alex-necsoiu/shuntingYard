# Go Calculator

A simple calculator implementation in Go, using the Shunting Yard Algorithm to parse and evaluate mathematical expressions.

## Implementation

The calculator is based on the Shunting Yard Algorithm, which converts infix expressions (normal mathematical notation) into postfix expressions (also known as Reverse Polish Notation or RPN). The postfix expression can then be evaluated easily using a stack.

## Features

- Supports basic arithmetic operations: addition, subtraction, multiplication, and division.
- Handles floating-point numbers and integer values.
- Respects operator precedence and associativity.
- Supports the use of parentheses to override the default precedence rules.

### Architecture

1. Define the data structure for operators: Create a struct `Operator` to store information about each operator, including its character representation, precedence, left associativity, and the function to evaluate the operator.

2. Initialize the operator map: Define a map called `operators` containing the supported operators (+, -, *, /) and their respective properties (precedence, associativity, and evaluation function).

3. Create utility functions: `higherPrecedence`, `isOperator`, and `tokenize` are utility functions that help with checking operator precedence, checking if a token is an operator, and tokenizing an input string, respectively.

4. Implement the Shunting Yard Algorithm in the `ShuntingYard` function: This function takes an input string and returns a postfix expression as a slice of strings. It uses a stack to handle operator precedence and associativity, as well as parentheses.

5. Implement the `EvaluatePostfix` function: This function takes a postfix expression (a slice of strings) and evaluates it using a stack, returning the result as a float64 value.

6. Create the `Calculate` function: This function combines the `ShuntingYard` and `EvaluatePostfix` functions to calculate the result of an infix expression. It first converts the input string into a postfix expression using the Shunting Yard Algorithm, then evaluates the postfix expression to obtain the final result.

## Usage

1. Build the Go application:

```shell
go build -o go-calculator
```

2. Run the calculator
   
```shell
./go-calculator
```

3. Run the test
   
```shell
go test
```

## 👥 Contributing

We welcome contributions in the form of bug reports, feature requests, or pull requests. For more information, please see CONTRIBUTING.md.

## 🔓 License
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

This project is licensed under the MIT License - see the LICENSE file for details.