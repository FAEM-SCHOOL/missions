package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var n, choose, elmt string

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Программа: Сортировка массива вставкой")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Сколько элементов будет в вашем массива? (1 <= n <= 1000)")
	fmt.Scan(&n)
	num ,_ := strconv.Atoi(n)
	fmt.Println("---------------------------------------------------------------------")

	fmt.Println("Заполнить массив рандомными значениями автоматически или вручную?")
	fmt.Println("1 - автоматически; 2 - вручную;")
	fmt.Scan(&choose)
	ch ,_ := strconv.Atoi(choose)
	fmt.Println("---------------------------------------------------------------------")

	array := make([]int, num)

	if ch == 1{
		for i := 0; i < num; i++{
			array[i]=rand.Intn(20000)-10000
		}
		fmt.Println(array, " ")

		for i := 0; i < num; i++{
			item := array[i]
			j := i
			for j > 0 && array[j - 1] > item{
				array[j] = array[j - 1]
				j--
				fmt.Println("---------------------------------------------------------------------")
				fmt.Println(array, " ")
			}
			array[j] = item
		}

		fmt.Println("---------------------------------------------------------------------")
		fmt.Println(array, " ")

	}else if ch == 2{
		for i := 0; i < num; i++{
			ret:
			fmt.Println("Диапазон чисел в массиве -10000 <= x <= 10000")
			fmt.Println("Введите " , i+1 , " число")
			fmt.Scan(&elmt)
			elm, _ :=strconv.Atoi(elmt)
				if elm < -10000 || elm > 10000{
					fmt.Println("Вы ввели число выходящее за заданный диапазон чисел!")
					goto ret
				}else if elm > -10001 || elm < 10001{
					array[i] = elm
				}
		}
		fmt.Println(array, " ")

		for i := 0; i < num; i++{
			item := array[i]
			j := i
			for j > 0 && array[j - 1] > item{
				array[j] = array[j - 1]
				j--
				fmt.Println("---------------------------------------------------------------------")
				fmt.Println(array, " ")
			}
			array[j] = item
		}

		fmt.Println("---------------------------------------------------------------------")
		fmt.Println(array, " ")
	}
}

