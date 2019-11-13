//22-2

//22-1

package main

import (
	"fmt"
	"sync"
)

func main() {
	//고루틴 동기화 고급
	//대기 그룹
	//실행 흐름 변경(고루틴이 종료될 때까지 대기 기능)
	//WaitGroup : Add(고루틴 추가), Done(작업 종료 알림), Wait(고루틴 종료 시까지 대기)
	//Add로 추가 된 고루틴 개수와 Done으로 종료되는 알림 횟수가 같아야 정확하게 동작한다. (Add == Done)

	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)

	for i := 0; i < 100; i++ {
		mutex.Lock()
		wg.Add(1)
		go func(n int) {
			fmt.Println("WaitGroup : ", n)
			wg.Done()
			mutex.Unlock()
		}(i)
	}
	wg.Wait() //모든 고루틴이 정상적으로 다 실행될 때까지 대기(Add == Done 횟수 같아야 함, 이 문법 사용하기 위해 Add와 Done사용)
	fmt.Println("WaitGroup End!")
}
