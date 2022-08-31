package main

func VocalKonsonan(input string) string {
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
