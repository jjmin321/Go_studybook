//1-7

package main

import "fmt"

func main() {
	//Data type : numeric
	//int
	//ex1
	//ASCII(EN)

	var char1 byte = 72
	var char2 byte = 0110
	var char3 byte = 0x48

	//ex2
	//unicode
	var char4 rune = 50556
	var char5 rune = 0142574
	var char6 rune = 0xC57c

	fmt.Printf("%c %c %c\n, %c %c %c", char1, char2, char3, char4, char5, char6)
}
