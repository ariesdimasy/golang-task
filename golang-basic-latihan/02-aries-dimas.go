package main

import (
	"fmt"
)

func vocalKonsonan(input string) string {
	var res string = "konsonan"

	var vocal string = "aiueoAIUEO"

	for i := 0; i < len(vocal); i++ {
		if input == string(vocal[i]) {
			res = "vocal"
			break
		}
	}

	return res
}

func main() {

	fmt.Print("Masukkan satu huruf : ")

	var input string

	fmt.Scanf("%s", &input)

	fmt.Println(vocalKonsonan(input))

}
