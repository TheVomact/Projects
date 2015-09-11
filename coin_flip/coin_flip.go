package main

import "fmt"
import "math/rand"
import "time"

func main() {
	var flip_num, heads, tails int

	fmt.Println("Input number of coins to flip.")
	_, err := fmt.Scanf("%d\n", &flip_num)
	if err != nil {
		fmt.Println("Invalid number.")
		return
	}

	/* Seeding with the same number has a bad habit of
	producing the same output on each run. */
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < flip_num; i++ {
		if rand.Intn(2)%2 == 0 {
			heads++
		} else {
			tails++
		}
	}

	fmt.Printf("Heads: %d, tails: %d\n", heads, tails)
}
