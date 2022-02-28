package main

import (
	"fmt"
	"strconv"
)

func main() {
	//tokens := []string{"2","1","+","3","*"}
	//g
	//fmt.Printf("%d\n", evalRPN(tokens))

	s := "-2 + 1"
	fmt.Printf("%v", calculate(s))
}

func evalRPN(tokens []string) int {

	var stack []int

	for _, token := range tokens {

		if token == "+" || token == "-" || token == "*" || token == "/" {

			y := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, calc(x, y, token))

		} else {

			num, _ := strconv.Atoi(token)

			stack = append(stack, num)
		}

	}
	return stack[0]
}

func calc(x int, y int, op string) int {
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	default:

	}
	return 0
}

func calculate(s string) int {
	var tokens []string
	var ops []byte

	n := len(s)
	needsZero := true
	for i := 0; i < n; i++ {

		switch s[i] {
		case ' ':
		case '(':
			needsZero = true
			ops = append(ops, '(')
		case ')':
			needsZero = false
			for len(ops) != 0 && ops[len(ops)-1] != '(' {
				tokens = append(tokens, string([]byte{ops[len(ops)-1]}))
				ops = ops[:len(ops)-1]
			}
			ops = ops[:len(ops)-1]
		case '+', '-', '*', '/':
			if (s[i] == '+' || s[i] == '-') && needsZero == true {
				tokens = append(tokens, "0")
			}
			for len(ops) != 0 && getRank(ops[len(ops)-1]) >= getRank(s[i]) {
				tokens = append(tokens, string([]byte{ops[len(ops)-1]}))
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, s[i])
			needsZero = false
		default:
			var num []byte
			for i < n && '0' <= s[i] && '9' >= s[i] {
				num = append(num, s[i])
				i++
			}
			i--
			tokens = append(tokens, string(num))
			needsZero = false
		}
	}
	for len(ops) != 0 {
		tokens = append(tokens, string([]byte{ops[len(ops)-1]}))
		ops = ops[:len(ops)-1]
	}

	fmt.Printf("%v", tokens)
	return evalRPN(tokens)
}

func getRank(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}
