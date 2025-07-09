package stackqueue

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, v := range tokens {
		if num, err := strconv.Atoi(v); err == nil {
			stack = append(stack, num)
		} else {
			num2, num1 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int
			switch v {
			case "+":
				res = num1 + num2
			case "-":
				res = num1 - num2
			case "*":
				res = num1 * num2
			case "/":
				res = num1 / num2
			}
			stack = append(stack, res)
		}
	}
	return stack[0]
}
