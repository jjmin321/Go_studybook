//5-1

package main

import "fmt"

func main() {
	//열거형
	//상수를 사용하는 일정한 규칙에따라 숫자를 계산 및 증가시키는 묶음

	//ex1
	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
		Jul
		Aug
		Sep
		Oct
		Nov
		Dec
	)

	fmt.Println("Jan : ", Jan)
	fmt.Println("Feb : ", Feb)
	fmt.Println("Mar : ", Mar)
	fmt.Println("Apr : ", Apr)
	fmt.Println("May : ", May)
	fmt.Println("Jun : ", Jun)
	fmt.Println("Jul : ", Jul)
	fmt.Println("Aug : ", Aug)
	fmt.Println("Sep : ", Sep)
	fmt.Println("Oct : ", Oct)
	fmt.Println("Nov : ", Nov)
	fmt.Println("Dec : ", Dec)

}
