// TODO study Go packages...
package utils

func Reverse(s string) (result string) {
	for _, c := range s {
		result = string(c) + result
	}
	return result
}
