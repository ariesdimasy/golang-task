package main

import (
	"fmt"
)

func isPalindrome(input string) bool {
	var res string = ""

	for i := len(input) - 1; i >= 0; i-- {
		res = res + string(input[i])
	}

	if res == input {
		return true
	} else {
		return false
	}
}

func main() {

	fmt.Print("Masukkan satu kata : ")

	var input string

	fmt.Scanf("%s", &input)

	fmt.Println((input))

}
