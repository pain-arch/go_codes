package main

import "fmt"

func getTwoVal(n1 int, n2 int) (int, int) {
	sum := n1 + n2

	mul := n1 * n2

	return sum, mul
}

func main() {

	a := 10
	b := 20

	sum, mul := getTwoVal(a,b)

	fmt.Println("this is the sum:",sum)
	fmt.Println("this is the mul:",mul)
}