package main

import (
	"fmt"
)

func calc_step(num, pow int) int {
	res := num
	for i := 1; i < pow; i++ {
		res *= num
	}
	return res
}

func main() {
	var num, pow int
	fmt.Print("input num and step separate space: ")
	_, _ = fmt.Scanf("%d %d", &num, &pow)

	var res = calc_step(num, pow)
	fmt.Println("resust: ", res)
}
