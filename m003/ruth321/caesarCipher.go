package main

import (
	"fmt"
)

func main() {

	var unenc string
	fmt.Print("Enter message (without spaces): ")
	_, _ = fmt.Scan(&unenc)

	var key byte
	fmt.Print("Enter key: ")
	_, _ = fmt.Scan(&key)

	enc := make([]byte, len(unenc))

	for i := 0; i < len(unenc); i++ {
		if unenc[i] > 64 && unenc[i] < 91 || unenc[i] > 96 && unenc[i] < 123 {
			if unenc[i] > 64 && unenc[i] < 91 {
				if unenc[i]+key > 90 {
					enc[i] = unenc[i] + key - 26
				} else {
					enc[i] = unenc[i] + key
				}
			}
			if unenc[i] > 96 && unenc[i] < 123 {
				if unenc[i]+key > 122 {
					enc[i] = unenc[i] + key - 26
				} else {
					enc[i] = unenc[i] + key
				}
			}
		} else {
			enc[i] = unenc[i]
		}
	}

	fmt.Print("Encrypted message: ")
	fmt.Println(string(enc))
}
