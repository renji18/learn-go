package main

import "fmt"

func calculator(a int, b int, operator string) (int, error) {
	if operator == "div" && b == 0 {
		return 0, fmt.Errorf("division by zero")
	}

	switch operator {
	case opAdd:
		return a + b, nil

	case opSub:
		return a - b, nil

	case opMul:
		return a * b, nil

	case opDiv:
		return a / b, nil

	default:
		return 0, fmt.Errorf("Unknown operator %s", operator)
	}
}
