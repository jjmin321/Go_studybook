package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strike int
	ball   int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	number := MakeNumbers() // make 3 random numbers

	cnt := 0
	for {
		input := InputNumbers() // input user's number

		result := CompareNumbers(number, input) // compare random numbers and user's number

		PrintResult(result)

		if IsGameEnd(result) {
			//game over
			break
		}
		cnt++
	}
	fmt.Printf("%d번만에 맞췄습니다", cnt)
	//game over. how many time?
}

func MakeNumbers() [3]int {
	var random [3]int
	for i := 0; i < 3; i++ {
		random[i] = rand.Intn(10)
	}
	fmt.Println(random)
	return random

}

func InputNumbers() [3]int {
	var rst [3]int
	for i := 0; i < 3; i++{
		fmt.Println("0~9 사이의 숫자 3개를 입력하세요.")
		var no int 
		_, err :	dasdf= fmt.Scanf("%d\n", &no)
	}

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력하세요.")
		var no int
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}

		success := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if duplicated {
				fmt.Println("숫자가 겹치지 않아야 합니다.")
				success = false
				break
			}

			if idx >= 3 {
				fmt.Println("3개보다 많은 숫자를 입력하셨습니다.")
				success = false
				break
			}

			rst[idx] = n
			idx++
		}
		if success && idx < 3 {
			fmt.Println("3개의 숫자를 입력하세요.")
			success = false
		}

		if !success {
			continue
		}
		break
	}
	rst[0], rst[2] = rst[2], rst[0]
	//fmt.Println(rst)
	return rst
}

func CompareNumbers(input, number [3]int) Result {
	strike := 0
	ball := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if input[i] == number[j] {
				if i == j {
					strike++
				} else {
					ball++
				}
			}
		}
	}
	return Result{strike, ball}
}

func PrintResult(result Result) {
	fmt.Printf("%ds, %db 입니다\n", result.strike, result.ball)
}

func IsGameEnd(result Result) bool {
	// 비교 결과가 3 스트라이크 인지 확인
	return result.strike == 3
}
