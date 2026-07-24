package main

import "fmt"

func main() {
	fmt.Println(isValid("([)]"))
}

func isValid(str string) bool {
	hash := map[rune]rune{')': '(', '}': '{', ']': '['}

	stack := []rune{}

	for _, s := range str {
		if val, found := hash[s]; found {
			if len(stack) > 0 && stack[len(stack)-1] == val {
				// pop item
				stack = stack[:len(stack)-1]
			}
		} else {
			// push item
			stack = append(stack, s)
		}
	}

	return len(stack) == 0
}


/*
	First we create a hash map to store the closing brackets as key and opening brackets as value.
	After that we loop over the string. If the single string item is found in the hash map, then we check if length of stack is > 0 && last item in stack == the value of the item found in hash map. If this condition is satisfies, then we pop the last item in the stack.
	If item is not found in stack, then we push the string item to the stack.
	After complete loop, if length of stack is 0 then the stack emptied completely, else not.
*/