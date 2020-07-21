package main

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)

func PrintArray(array []int)  {
	for e := 0; e < len(array) - 1; e++{
		fmt.Print(array[e], " ")
	}
	fmt.Println(array[len(array) - 1])
}

func main() {
	var n int
	fmt.Scan(&n)

	array := make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := strings.Split(scanner.Text(), " ")

	for i := 0; i < len(input); i++ {
		array[i], _ = strconv.Atoi(input[i])

	}

	for i := 1; i < n; i++ {
		temp := array[i]
		j := i
		for j > 0 && array[j - 1] > temp {
			array[j] = array[j - 1]
			j--
		}
		array[j] = temp
	}
	PrintArray(array)
}

