//15-8

package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) CalculateD() {
	a.balance = a.balance + (a.balance * a.interest)
}

/*
func CalculateD (a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}
*/

func (a *Account) CalculateP() {
	a.balance = a.balance + (a.balance * a.interest)
}

/*
func CalculateP (a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}
*/
func main() {
	//ex1
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-902", 12000000, 0.015}

	kim.CalculateD()
	// CalculateD(kim)
	lee.CalculateP() //리시버 메소드 사용시 &를 사용하지 않더라도 자동으로 포인터형으로 보내짐
	// CalculateP(&lee) 함수 사용시 강제적으로 &를 사용해 포인터형으로 보내줘야함.

	fmt.Println(int(kim.balance))
	fmt.Println(int(lee.balance))
}
