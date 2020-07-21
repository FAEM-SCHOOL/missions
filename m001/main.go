package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	var str string
	fmt.Scan(&str)

	var k int
	fmt.Scan(&k)

	k = k % 26

	var result string = ""
	for i := 0; i < len(str); i++ {

		var sym rune = rune(int(str[i]))

		if sym >= 65 && sym <= 90 {
			if int(sym)+k > 90 {
				ost := 90 - int(sym)
				sym = rune(65 + (k - ost - 1))
			} else {
				sym = rune(int(sym) + k)
			}
		} else {
			if sym >= 97 && sym <= 122 {
				if int(sym)+k > 122 {
					ost := 122 - int(sym)
					sym = rune(97 + (k - ost - 1))
				} else {
					sym = rune(int(sym) + k)
				}
			}
		}

		result += string(sym)
	}

	fmt.Println(result)
}
