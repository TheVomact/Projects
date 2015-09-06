package main

import "fmt"

func main() {
	var num int
	var execution_result int

	execution_result, _ = fmt.Scanln(&num)
	if execution_result != 1 {
		return
	}

	fmt.Println(calc_fibb(num))
}

func calc_fibb(iterations int) int {
	var current, next int = 0, 1

	// We've already done one iteration...
	iterations -= 1

	for j := 0; j < iterations; j++ {
		next, current = current+next, next
	}
	return current
}
