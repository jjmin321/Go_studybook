package main

import "fmt"

func main() {
	//배열 순회

	//ex1 use for loop
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(i, arr1[i])
	}
	//ex2 use for range loop
	arr2 := [5]int{7, 77, 777, 7777, 77777}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
}
