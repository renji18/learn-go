package main

import (
	"fmt"
)

func one() {
	reverseString("नमस्ते")
	reverseString("你好")
	checkPalindrome("नमन")
}

func reverseString(str string) {
	runeArr := []rune(str)

	// for _, r := range str {
	// 	runeArr = append(runeArr, r)
	// }

	// slices.Reverse(runeArr)

	// for left, right := 0, len(runeArr)-1; left < right; left++ {
	// 	leftRef := runeArr[left]
	// 	runeArr[left] = runeArr[right]
	// 	runeArr[right] = leftRef
	// 	right--
	// }

	for left, right := 0, len(runeArr)-1; left < right; left, right = left+1, right-1 {
		runeArr[left], runeArr[right] = runeArr[right], runeArr[left]
	}

	// for _, s := range runeArr {
	// 	ans = ans + string(s)
	// }

	ans := string(runeArr)

	fmt.Println(ans)
}

func checkPalindrome(str string) {
	runeArr := []rune(str)

	right := len(runeArr) - 1

	for left := 0; left < right; left++ {
		if runeArr[left] != runeArr[right] {
			fmt.Println("Not Palindrome")
			return
		}
		right--
	}

	fmt.Println("Palindrome")
}
