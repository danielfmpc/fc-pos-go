package main

import "fmt"

func main() {
	soma := sum(3, 5)
	fmt.Println(soma)
}

func sumVariadic(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func sum(a, b int) int {
	return a + b
}
