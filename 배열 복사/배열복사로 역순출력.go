package main

import "fmt"

func main() { //2n의 반복으로 역순출력
	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}
	fmt.Println("arr:", arr)

	for i := 0; i < len(arr); i++ {
		clone[i] = arr[len(arr)-1-i]
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = clone[i]
	}

	fmt.Println("clone:", clone)
	fmt.Println("arr:", arr)

}
