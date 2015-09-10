package main

import "fmt"

var units_map = map[string]float64{"kHz": 1000, "daHz": 100, "hHz": 10, "Hz": 1}
var units = [4]string{"kHz", "daHz", "hHz", "Hz"}

func main() {
	var num float64
	var unit, to_unit string

	fmt.Println("Enter the value and its unit freq.")
	_, err := fmt.Scanf("%f %s\n", &num, &unit)
	if err != nil || !sanitizeInput(unit) {
		fmt.Println("Invalid input.")
		return
	}

	fmt.Println("Enter the unit to which to convert the amount.")
	_, err = fmt.Scanln(&to_unit)
	if err != nil || !sanitizeInput(to_unit) {
		fmt.Println("Invalid input.")
		return
	}

	convertedValue := convertUnit(num, unit, to_unit)
	fmt.Printf("Converted value is %.2f %s.\n", convertedValue, to_unit)
}

func sanitizeInput(unit string) bool {
	for _, u := range units {
		if u == unit {
			return true
		}
	}
	return false
}

func convertUnit(num float64, unit string, to_unit string) float64 {
	var multiplier float64
	var divisor float64
	var base_value float64

	multiplier = units_map[unit]
	base_value = multiplier * num

	divisor = units_map[to_unit]
	return base_value / divisor
}
