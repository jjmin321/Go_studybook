//22-3

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//고루틴 동기화 고급
	//원자성 사용 -> 기능적으로 분할 불가능한 완전 보증된 일련의 조작, 모두 성공하거나 모두 실패
	//모든 조작이 완료될 때까지 다른 프로세스 개입 불가
	//sync/atomic에서 원자적 연산자 제공
	//https://golang.org/pkg/sync/atomic
	//주로 공용 변수에 관한 계산 사용

	//원자성 사용하지 않을 경우
	//ex1
	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0

	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)

	for i := 0; i < 5000; i++ {
		//mutex.Lock()
		wg.Add(1)
		go func(n int) {
			cnt++
			wg.Done()
			//mutex.Unlock()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		// mutex.Lock()
		wg.Add(1)
		go func(n int) {
			cnt--
			wg.Done()
			// mutex.Unlock()
		}(i)
	}
	wg.Wait() //모든 고루틴이 정상적으로 다 실행될 때까지 대기(Add(7000) == Done(7000) 횟수 같아야 함, 이 문법 사용하기 위해 Add와 Done사용)
	fmt.Println("WaitGroup End! >>>>>", cnt)
}
