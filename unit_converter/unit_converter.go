package main

import "fmt"

func main() {
	var num, convertedValue int
	var unit, to_unit string

	fmt.Println("Enter the value and its unit mass.")
	fmt.Scanf("%d %s\n", &num, &unit)
	if !sanitizeInput(unit) {
		fmt.Println("Invalid input.")
		return
	}
	fmt.Println("Enter the unit to which to convert the amount.")
	fmt.Scanln(&to_unit)
	if !sanitizeInput(to_unit) {
		fmt.Println("Invalid input.")
		return
	}

	convertedValue = convertUnit(num, unit, to_unit)
	fmt.Printf("Converted value is %d %s.\n", convertedValue, to_unit)
}

func sanitizeInput(unit string) bool {
	return true
}

func convertUnit(num int, unit string, to_unit string) int {
	return 0
}
