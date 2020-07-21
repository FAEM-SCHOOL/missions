package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите пароль:")
	var password string
	fmt.Fscan(os.Stdin, &password)
	fmt.Println(checkPassword(password))
}

// checkPassword принимает пароль и возвращает минимальное число символов, которые необходимо добавить,
// чтобы пароль считался сильным.
//
// Cильный пароль имеет:
//	- длину > 5
//	- хотя бы 1 цифру
//	- хотя бы 1 символ нижнего регистра
//	- хотя бы 1 символ верхнего регистра
//	- хотя бы 1 спецсимвол: !@#$%^&*()-+
func checkPassword(pwd string) int {
	const numbers = "0123456789"
	const lowerCase = "abcdefghijklmnopqrstuvwxyz"
	const upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const specialCharacters = "!@#$%^&*()-+"
	count := 0
	if !strings.ContainsAny(pwd, numbers) {
		count++
	}
	if !strings.ContainsAny(pwd, lowerCase) {
		count++
	}
	if !strings.ContainsAny(pwd, upperCase) {
		count++
	}
	if !strings.ContainsAny(pwd, specialCharacters) {
		count++
	}
	if remain := 6 - len(pwd); remain > count {
		count = remain
	}
	return count
}
