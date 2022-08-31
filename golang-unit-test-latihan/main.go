package main

import (
	"fmt"
)

func main() {

	fmt.Print("Masukkan satu kata : ")

	var input1 string

	fmt.Scanf("%s", &input1)

	fmt.Println(IsPalindrome(input1))

	fmt.Print("Masukkan satu huruf : ")

	var input2 string

	fmt.Scanf("%s", &input2)

	fmt.Println(VocalKonsonan(input2))

}
