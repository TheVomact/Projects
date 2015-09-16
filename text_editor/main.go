package main

import "fmt"
import "github.com/nsf/termbox-go"

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Printf("termbox.Init() returned an error: \n%s\n", err)
		return
	}

	termbox.Close()
}
