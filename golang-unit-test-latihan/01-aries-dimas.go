package main

func IsPalindrome(input string) bool {
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
