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
	fmt.Println("studyResult", studyResult)

	fmt.Println(reflect.TypeOf(s).Field(1).Name)
	sf, _ := reflect.TypeOf(s).FieldByName("name")

	fmt.Println("sf.Index", sf.Index)
	fmt.Println("sf.Type", sf.Type)
	fmt.Println("sf.Name", sf.Name)
	fmt.Println("reflect.ValueOf(sf.Index)---", reflect.ValueOf(sf))
	fmt.Println(reflect.TypeOf(sf.Index))
	fmt.Println(sf.Name)

	fmt.Println(reflect.ValueOf(s).FieldByIndex(sf.Index))

	fmt.Println(reflect.ValueOf(s.name))
}
