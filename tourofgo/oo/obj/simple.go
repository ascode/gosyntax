package main

import (
	"fmt"
	"reflect"
)

type student struct {
	id int
	name string
}

type Curriculum struct {
	name string
}

func (s student) Study(c Curriculum) (string, error) {
	return c.name, nil
}

func main() {
	s := student{id:1, name: "jinfei"}
	studyResult, _ := s.Study(Curriculum{name: "go课程"})
	fmt.Println(studyResult)

	fmt.Println(reflect.TypeOf(s).Field(1).Name)
	fmt.Println(reflect.ValueOf(s.name))
}
