package main //맨앞부터 1개씩없애고 없앤 수 출력하는거

import (
	"fmt"
)

func RemoveA(a []int) ([]int, int) {
	return a[1:], a[0]	//a[1]부터 끝까지 반환, 없어진 a[0]을 출력 
}	//a[0]을 없애는 게 아니라 없앤 것 처럼 보이는 것 뿐이다.

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 5; i++ {
		var back int
		a, back = RemoveA(a) //a의 첫 번째 수를 첫번째수 +1로 초기화
		fmt.Printf("%d\n %d\n", a, back)
	}
}
