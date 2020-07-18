package main

import "fmt"

func main() {

	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	c := ai1(a, b)
	fmt.Println(c)
}

func ai1(num, step int) int {
	var otv int = 1
	for i := 0; i < step; i++ {
		otv *= num
	}
	return otv
}