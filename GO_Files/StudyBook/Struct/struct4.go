package main

import "fmt"

func main() {
	//구조체 익명선언 및 활용

	//ex1
	//type 타입 - 익명 구조체
	car1 := struct{ name, color string }{"520d", "red"}

	fmt.Println(car1)
	fmt.Printf("%#v\n", car1)

	//ex2
	cars := []struct{ name, color string }{{"520d", "red"}, {"530i", "blue"}, {"528i", "black"}}
	for _, b := range cars {
		fmt.Printf("(%s %s) ----- (%#v)\n", b.name, b.color, b)
	}

}
