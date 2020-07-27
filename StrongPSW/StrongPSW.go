package main

import (
	"fmt"
)

var pas string
var warn int

func main() {
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Println("Программа: 'Сложный пароль'")
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Println("Сложный пароль должен: \n"+"1) Содержать не менее 6 символов\n"+"2)Содержать не менее 1 цифры\n"+
		"3)Содержать латинские буквы \n"+"4)Содержать не менее 1 буквы нижнего регистра \n"+
		"5)Содержать не менее 1 буквы верхнего регистра \n"+"6)Содержать не менее 1 специального символа")
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Println("FAQ")
	fmt.Println("Цифры: 123456789 \n"+"Специальные символы: !@#$%^&*()-+ \n"+
		"Буквы верхнего регистра: ABCDEFGHIJKLMNOPQRSTUVWXYZ \n"+"Буквы нижнего регистра: abcdefghijklmnopqrstuvwxyz")
	fmt.Println("------------------------------------------------------------------------------------")

	fmt.Println("Введите ваш пароль: ")
	fmt.Scan(&pas)

	for i := 0; i < 1; i++ {
		if CheckDnReg(pas) && CheckNumPsw(pas) && CheckSizePsw(pas) && CheckSpecSym(pas) && CheckUpReg(pas) {
			fmt.Println("Сложный пароль.")
			break
		}else{
			if !CheckUpReg(pas){
				warn++
				fmt.Println(warn, ")Ваш пароль не содержит буквы ВЕРХНЕГО регистра!")
			}
			if !CheckDnReg(pas){
				warn++
				fmt.Println(warn, ")Ваш пароль не содержит буквы НИЖНЕГО регистра!")
			}
			if !CheckSpecSym(pas){
				warn++
				fmt.Println(warn, ")Ваш пароль не содержит СПЕЦИАЛЬНЫЕ символы!")
			}
			if !CheckNumPsw(pas){
				warn++
				fmt.Println(warn, ")Ваш пароль не содержит ЦИФРЫ!")
			}
			if !CheckSizePsw(pas){
				warn++
				fmt.Println(warn, ")Ваш пароль НЕ ДОСТАТОЧНО ДЛИННЫЙ!")
			}
		}
	}
}

func CheckSizePsw(psw string) (a bool) {
	if len(psw) >= 6{
		a = true
		return a
	}
	a = false
	return a
}

func CheckNumPsw(psw string) (a bool)  {
	for i := 0; i < len(psw); i++{
		if psw[i]>=48 && psw[i]<=57{
			a = true
			return a
		}
	}
	a = false
	return a
}

func CheckUpReg(psw string) (a bool) {
	for i := 0; i < len(psw); i++{
		if psw[i]>=65 && psw[i]<=90{
			a = true
			return a
		}
	}
	a = false
	return a
}

func CheckDnReg(psw string) (a bool) {
	for i := 0; i < len(psw); i++{
		if psw[i]>=97 && psw[i]<=122{
			a = true
			return a
		}
	}
	a = false
	return a
}

func CheckSpecSym(psw string) (a bool) {
	for i := 0; i < len(psw); i++{
		if (psw[i]>=33 && psw[i]<=47) || (psw[i]>=58 && psw[i]<=64) ||
			(psw[i]>=91 && psw[i]<=96) || (psw[i]>=123 && psw[i]<=126){
			a = true
			return a
		}
	}
	a = false
	return a
}