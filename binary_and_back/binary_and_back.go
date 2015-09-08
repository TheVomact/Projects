package main

import "fmt"
import "strconv"
import "os"
import "github.com/TheVomact/Projects/utils"

func main() {
	var num int
	var result string

	if len(os.Args) > 3 || len(os.Args) == 1 {
		fmt.Println("Invalid argument number.")
	} else if os.Args[1] == "b" {
		fmt.Scanln(&num)
		if !isNumberBinary(num) {
			fmt.Println("Invalid number format.")
			return
		}
		result = BinaryToDecimal(num)
	} else if os.Args[1] == "d" {
		fmt.Scanln(&num)
		result = DecimalToBinary(num)
	} else {
		fmt.Println(`Invalid argument to program. Program accepts 'b' - binary -> decimal
		or 'd' - decimal -> binary`)
	}

	fmt.Println(result)
}

func isNumberBinary(num int) bool {
	str := strconv.Itoa(num)
	for _, r := range str {
		char := string(r)
		if !(char == "1" || char == "0") {
			return false
		}
	}
	return true
}

func BinaryToDecimal(num int) string {
	var decimal_repr int
	multiplier := 1
	str := strconv.Itoa(num)
	str = utils.Reverse(str)
	for _, r := range str {
		n, _ := strconv.Atoi(string(r))
		decimal_repr += n * multiplier
		multiplier *= 2
	}

	return strconv.Itoa(decimal_repr)
}

func DecimalToBinary(num int) string {
	var binary_repr string
	for {
		result := strconv.Itoa(num % 2)
		num = num / 2
		binary_repr = result + binary_repr
		if num == 0 {
			break
		}
	}

	return binary_repr
}
