package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reversed, str string
	string_slice := make([]string, 3)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string.")
	for scanner.Scan() {
		string_slice = append(string_slice, scanner.Text())
	}
	str = strings.Join(string_slice, " ")
	for _, ch := range str {
		reversed = string(ch) + reversed
	}

	fmt.Println(reversed)
}
