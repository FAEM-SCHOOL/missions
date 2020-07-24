package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var arr []int

func defenitio() {
	var def int
	var scan int
	var command string
	for {
		fmt.Print("Введите размер массива, размер должен быть больше 2 и не превышать 100: ")
		fmt.Scanln(&command)
		kek, err := strconv.ParseInt(command, 10, 32)
		if err == nil && kek > 2 && kek < 100{
			def = int(kek)
			break
		}
		fmt.Println("Введите корректный размер!")
	}
	fmt.Println("1 - ручное заполнение массива")
	fmt.Println("2 - рандомное заполнение массива")
	fmt.Scan(&scan)
	if scan == 1 {
		for i := 0; i < (def); i++ {
			fmt.Print("Введите элемент массива arr[", i, "] : ")
			fmt.Scan(&scan)
			arr = append(arr, scan)
		}
	} else if scan == 2 {
		for i := 0; i < (def); i++ {
			a := -10000
			b := 10000
			rand.Seed(time.Now().UnixNano())
			var rand = a + rand.Intn(b-a+1)
			arr = append(arr, rand)
		}
	} else {
		fmt.Println("Введите верную команду!")
		defenitio()
	}
	fmt.Println("Вывод неотсортированного массива:")
	fmt.Println(arr)
}

func sorting() {
	for i := 0; i < len(arr); i++ {
		o := i
		for j := i - 1; j > -1; j-- {
			if arr[o] < arr[j] {
				arr[j], arr[o] = arr[o], arr[j]
				o--
			}
		}
	}
	fmt.Println("Вывод отсортированного массива:")
	fmt.Println(arr)
}

func main() {
	defenitio()
	sorting()
}
