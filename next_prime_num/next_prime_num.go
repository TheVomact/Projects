package main

import "fmt"
import "math"

func main() {
	current_prime := 1.
	fmt.Println("Current prime is:\n1")

	for {
		var input string
		fmt.Println("Fetch the next prime number? y/N?")
		fmt.Scanln(&input)

		if input != "y" {
			return
		}

		current_prime = fetchNextPrimeNum(current_prime)
		fmt.Println("Next prime is:", current_prime)
	}
}

/* Fun fact. In Go modulo operator is not defined for float64.
Only for integers (at least I think so). Use math package. */
func fetchNextPrimeNum(prime float64) float64 {
	for {
		prime += 1.
		var i float64

		for i = 2.0; i <= prime; i++ {
			if math.Mod(prime, i) == 0 {
				break
			}
		}

		if prime == i {
			return prime
		}
	}
}
