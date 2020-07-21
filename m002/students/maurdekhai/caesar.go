package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Print("Введите текст (латиница): ")
	var plaintext string
	fmt.Fscan(os.Stdin, &plaintext)

	fmt.Print("Введите сдвиг: ")
	var rot int
	fmt.Fscan(os.Stdin, &rot)

	ciphertext := rotateString(plaintext, rot)
	fmt.Println(ciphertext)
}

// rotateString шифрует строку s шифром Цезаря по ключу rot
func rotateString(s string, rot int) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, char := range s {
		b.WriteRune(rotateChar(char, rot))
	}

	return b.String()
}

// rotateChar "сдвигает" буквенный символ char на rot позиций.
// Символ не изменяется, если он не принадлежит английскому алфавиту.
func rotateChar(char rune, rot int) rune {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	// приводим к нижнему регистру, чтобы не создавать
	// отдельный алфавит для верхнего регистра
	wasUpperCase := unicode.IsUpper(char)
	char = unicode.ToLower(char)
	if index := strings.IndexRune(alphabet, char); index > -1 {
		index += rot
		// возвращаем индекс в границы алфавита
		length := len(alphabet)
		if index < 0 {
			index += length
		}
		index = index % length

		char = []rune(alphabet)[index]
		if wasUpperCase {
			char = unicode.ToUpper(char)
		}
	}
	return char
}
