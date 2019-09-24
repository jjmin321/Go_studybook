package main

import "fmt"

func main() { //n/2번의 반복으로 역순출력
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Print(arr)
}
