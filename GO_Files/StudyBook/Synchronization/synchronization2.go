//21-2

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//구조체 선언(공유 데이터)
type count struct {
	num   int
	mutex sync.Mutex
}

func (c *count) increment() {
	//공유 데이터 수정 전 뮤텍스로 보호 //한 명 들어가고 문 잠궈놓고
	c.mutex.Lock()
	c.num++
	//공유 데이터 수정 후 뮤텍스로 보호 해제 //작업끝나면 문 열고 나가고가 반복됨.
	c.mutex.Unlock()
}

func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	//고루틴 동기화 예제
	//실행 흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장 중요
	//뮤텍스(Mutex) : 여러 고루틴에서 작업하는 공유데이터 보호
	//sync.Mutex 선언 후 Lock, Unlock 두 개로 사용

	//동기화 사용한 경우
	//시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)
	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() //cpu 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result()
}
