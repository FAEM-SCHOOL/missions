package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter unsorted array: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s := scan.Text()
	a := strToIntSlice(s)

	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			b := a[i]
			g := i
			for ; g > 0 && b < a[g-1]; g-- {
				a[g] = a[g-1]
			}
			a[g] = b
		}
	}

	fmt.Print("Sorted array: ")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()

}

func strToIntSlice(s string) []int {
	c := strings.Fields(s)
	a := make([]int, len(c))
	for i := 0; i < len(c); i++ {
		a[i], _ = strconv.Atoi(c[i])
	}
	return a
}
