//18-4

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//고루틴(Goroutine)
	//클로저 사용 예제

	//ex1
	runtime.GOMAXPROCS(2)

	s := "Goroutine Closure : "

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) //반복문 클로저는 일반적으로 즉시 실행하지만 고루틴 클로저는 반복문이 종료된 후 우다다다다 실행
	}
	time.Sleep(1 * time.Second)
}
