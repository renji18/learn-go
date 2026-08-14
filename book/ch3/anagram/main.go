package main

import (
	"flag"
	"fmt"
	"os"
)

var str1 = flag.String("s1", "", "string to check anagram of")
var str2 = flag.String("s2", "", "string to check anagram of")

// func main() {
// 	flag.Parse()

// 	fmt.Println(*str, "str")

// 	strBytes := []byte(*str)

// 	for i := range len(strBytes) / 2 {
// 		if strBytes[i] != strBytes[len(strBytes)-1-i] {
// 			fmt.Println("Not anagram")
// 			os.Exit(1)
// 		}
// 	}

// 	fmt.Println("Anagram")
// }

func main() {
	flag.Parse()

	str1Bytes := []byte(*str1)
	str2Bytes := []byte(*str2)

	str1Len := len(str1Bytes)

	if str1Len != len(str2Bytes) {
		fmt.Println("Not Anagram")
		os.Exit(2)
	}

	for i := range str1Len {
		if str1Bytes[i] != str2Bytes[str1Len-i-1] {
			fmt.Println("Not Anagram")
			os.Exit(1)
		}
	}

	fmt.Println("Anagram")
}
