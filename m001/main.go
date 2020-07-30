package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)

	var pass string
	fmt.Scan(&pass)

	var result int

	var numbers bool = false
	var lower_case bool = false
	var upper_case bool = false
	var special_char bool = false

	for i := 0; i < len(pass); i++ {

		var sym rune = rune(pass[i])

		if sym >= '0' && sym <= '9' { numbers = true }
		if sym >= 'a' && sym <= 'z' { lower_case = true }
		if sym >= 'A' && sym <= 'Z' { upper_case = true }
		if sym == '!' || sym == '@' || sym == '-'  || sym == '^'|| (sym >= '#' && sym <= '&') || (sym >= '(' && sym <= '+'){
			special_char = true
		}
	}

	if(!numbers) { result++ }
	if(!lower_case) { result++ }
	if(!upper_case) { result++ }
	if(!special_char) { result++ }

	if(len(pass) + result < 6) { result += 6 - (len(pass) + result) }

	fmt.Println(result)
}

