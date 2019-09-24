package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintScore() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputScore(class string, grade string) {
	s.class = class
	s.grade = grade
}
func main() {
	var s Student = Student{name: "Tucker", age: 23, class: "수학", grade: "A+"}

	s.InputScore("과학", "C")
	s.PrintScore()
}
