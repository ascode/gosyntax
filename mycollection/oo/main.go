package main

import "fmt"

func main() {

	stu := Student{
		Name:"jifnei",
		Age:10,
	}
	stu.Grow(11)
	fmt.Println(stu.Age)

	soft := Software{
		stds:[]Student{stu},
	}

	soft.Add(Student{
		Name:"j1",
		Age:13,
	})
	fmt.Println(soft)

}

func (s Student) Grow(age int64) {
	s.Age = age
}

func (s *Software) Add(std Student){
	s.stds = append(s.stds, std)
}

type Student struct {
	Age int64
	Name string
}

type Software struct {
	stds []Student
}
