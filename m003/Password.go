package main

import (
	"fmt"
)

var command string
var number bool
var lower bool
var upper bool
var special bool

//Проверка пароля на наличие необходимых символов.
func chek_length() {
	if len(command) < 6 {
		fmt.Println("Длинв пароля слишком мала!!!")
		return
	}
}
func chek_number() {
	for i := 0; i < len(command); i++ {
		numbers := "0123456789"
		for j := 0; j < len(numbers); j++ {
			if command[i] == numbers[j] {
				number = true
				break
			}
		}
	}

	if !number {
		fmt.Println("НЕТ ЦИФР  В ПАРОЛЕ!!!")
	}
}
func chek_lower() {
	for i := 0; i < len(command); i++ {
		lower_case := "abcdefghijklmnopqrstuvwxyz"
		for j := 0; j < len(lower_case); j++ {
			if command[i] == lower_case[j] {
				lower = true
				break
			}
		}
	}
	if !lower {
		fmt.Println("НЕТ ПРОПИСНЫХ БУКВ  В ПАРОЛЕ!!!")
	}
}
func chek_upper() {
	for i := 0; i < len(command); i++ {
		upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for j := 0; j < len(upper_case); j++ {
			if command[i] == upper_case[j] {
				upper = true
				break
			}
		}
	}
	if !upper {
		fmt.Println("НЕТ ЗАГЛАВНЫХ БУКВ В ПАРОЛЕ!!!")
	}
}
func chek_special() {
	for i := 0; i < len(command); i++ {
		special_characters := "!@#$%^&*()-+"
		for j := 0; j < len(special_characters); j++ {
			if command[i] == special_characters[j] {
				special = true
				break
			}
		}
	}
	if !special {
		fmt.Println("НЕТ СПЕЦ СИМВОЛОВ В ПАРОЛЕ!!!")
	}
}

func main() {
	fmt.Print("Введите пароль для проверки :")
	fmt.Scan(&command)
	chek_length()
	chek_number()
	chek_lower()
	chek_upper()
	chek_special()
	if special && lower && number && upper {
		fmt.Println("Ваш пароль сильный!!!")
	}
}
