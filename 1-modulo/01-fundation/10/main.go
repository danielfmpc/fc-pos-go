package main

import "fmt"

func main() {
	soma := sumVariadic(1, 2, 3, 4, 5)

	total := func() int {
		return sumVariadic(1, 2, 3, 4, 5) * 2
	}()

	fmt.Println(total)
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
