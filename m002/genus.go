package main

import (
	"fmt"
	"math/rand"
	"time"
)
var arr []int
//Функция заполнения массива, выбор ручного и автоматического режима заполнения.
func defenitio(){
	var def int
	var scan int
	fmt.Print("Введите размер массива, его размер не должен превышать 100 и быть менее 2:")
	fmt.Scan(&def)
	if def<2 || def>100{
		fmt.Println()
		main()
	}
	fmt.Println("1 - ручное заполнение массива")
	fmt.Println("2 - рандомное заполнение массива")
	fmt.Scan(&scan)
	if scan==1 {
		for i:=0;i<(def);i++{
			fmt.Print("Введите элемент массива arr[",i,"] :")
			fmt.Scan(&scan)
			arr=append(arr,scan)
		}
	}else {
		for i:=0;i<(def);i++{
			a:=-10000
			b:=10000
			rand.Seed(time.Now().UnixNano())
			var rand= a + rand.Intn(b-a+1)
			arr=append(arr,rand)
	}}
	fmt.Println("Вывод неотсортированного массива:")
	fmt.Println(arr)
}
//Функция сортировки вставкой.
func sorting(){
	for i:=0;i< len(arr);i++{
		o:=i
		for j := i-1; j > -1 ;j--{
			if arr[o]<arr[j]{
				arr[j],arr[o]=arr[o],arr[j]
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
