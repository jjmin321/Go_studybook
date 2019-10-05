package main

import "fmt"

func main() {
	//배열 복사
	//값 복사 확인 important important important important important

	//ex2
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1

	fmt.Println(arr1, &arr1)
	fmt.Println(arr2, &arr2)

	const (
		a = iota * 5
		b
		c
		d
		e
	)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = e
	}

	fmt.Printf("%v %p\n", arr1, &arr1)
	fmt.Printf("%d %p\n", arr2, &arr2)
}
