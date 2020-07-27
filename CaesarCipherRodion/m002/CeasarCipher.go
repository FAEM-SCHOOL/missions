package main

import (
	"fmt"
	"strconv"
)

var fmes,mes, key string

func main() {
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Println("Программа: 'Шифр Цезаря'")
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Println("Шифр Цезаря — это вид шифра подстановки, \n" +
		"в котором каждый символ в открытом тексте заменяется символом, \n" +
		"находящимся на некотором постоянном числе правее него в алфавите. ")
	fmt.Println("------------------------------------------------------------------------------------")

	fmt.Println("Введите сообщение:")
	fmt.Scan(&mes)
	fmt.Println("Введите размер сдвига:")
	fmt.Scan(&key)
	shift,_ :=strconv.Atoi(key)

	runes:=[]rune(mes)

	var result []int

	for i := 0; i < len(runes); i++{
		result = append(result, int(runes[i]))

		AsciiId := (result[i] >= 65) && (result[i] <= 90) || (result[i] >= 97) && (result[i] <= 122)
		if AsciiId {

			result[i] = shift + result[i]
			//Если мы выходим за пределы 90 и 122, то программа возвращается к началу ASCII
			if result[i] > 90{
				rem := 90 - result[i]
				result[i] = 64 - rem
			}

			if result[i] > 122{
				rem := 122 - result[i]
				result[i] = 96 - rem
			}
		}
		fmes += string(result[i])
	}

	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Println("ASCII код: " , result)
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Println("Зашифрованное сообщение: " , fmes)
	fmt.Println("------------------------------------------------------------------------------------")
}


