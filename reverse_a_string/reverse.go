package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reversed, str string
	fmt.Println("Enter a string.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()

	for _, ch := range str {
		reversed = string(ch) + reversed
	}

	fmt.Println(reversed)
}
