//10-1

package main

import "fmt"

func main() {
	//포인터
	//GO : 포인터 지원(C)
	//변수의 지역성, 연속된 메모리 참조 ..., 힙, 스택... , 주관적인 생각 : 포인터는 폭력성이 있다. 메모리에 접근해서 뻇어오기 때문에.
	//파이썬, 자바(JRE) -> 컴파일러, 인터프리터
	//포인터지원 X(파이썬, C#, JAVA 등)
	//주소의 값은 직접 변경이 불가능(잘못된 코딩으로 인한 버그 방지)
	//*(애스터리스크)로 사용
	//nil로 초기화 (nil == 0)

	//ex1
	var a *int            //방법1 (nil로 초기화됨)
	var b *int = new(int) //방법2 (주소값을 가짐) 객관적으로 더 선호

	fmt.Println(a) // <nil>
	fmt.Println(b) // 0xc0000120a8와 같은 메모리주소

	i := 7

	a = &i //&주소값 전달
	b = &i //&주소값 전달

	fmt.Println(a, &i) // a -> i의 메모리주소 , &i -> i의 메모리주소
	fmt.Println(&a)    // &a -> a의 메모리주소
	fmt.Println(*a)    // *a -> i의 메모리주소의 값 (a는 포인터형 정수이므로 메모리주소를 가짐)
	fmt.Println()
	fmt.Println(b, &i) // b -> i의 메모리주소 , &i -> i의 메모리주소
	fmt.Println(&b)    // &b -> b의 메모리주소
	fmt.Println(*b)    // *b -> i의 메모리주소의 값 (b는 포인터형 정수임으로 메모리주소를 가짐)

	var c = &i
	d := &i

	*d = 10            //d는 i의 주소를 가지고 있으므로 값을 참조하기 위해 *를 붙혀야 함.
	fmt.Println(c, &i) // c -> i의 메모리주소 , &i -> i의 메모리주소
	fmt.Println(&c)    // &c -> a의 메모리주소
	fmt.Println(*c)    // *c -> i의 메모리주소의 값 (c는 포인터형 정수이므로 메모리주소를 가짐)

	fmt.Println()

	fmt.Println(d, &i) // d -> i의 메모리주소 , &i -> i의 메모리주소
	fmt.Println(&d)    // &d -> d의 메모리주소
	fmt.Println(*d)    // *d -> i의 메모리주소의 값 (d는 포인터형 정수임으로 메모리주소를 가짐)

}
