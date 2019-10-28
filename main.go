package main

import (
	"fmt"
	"github.com/ascode/rset/collection"
	"github.com/ascode/rset/test_data"
)

func main() {
	// 新建一个RSet
	rset := collection.NewSet(test_data.StudentsGood, test_data.StudentsBad)

	printSet(rset)
	fmt.Println("---------------")
	rset.SortDescDowngradeBy("Id") 	// 集合按照排序字段顺次降级排序
	printSet(rset)
	fmt.Println("---------------")
	rset.SortAscDowngradeBy("Id") 	// 集合按照排序字段顺次升级排序
	printSet(rset)

	fmt.Println("")
	fmt.Println("----------------------------------------------")


	s := rset.SkipReturn(1).LimitReturn(2) // 留下第1个开始连续2个元素
	printSet(&s)

	fmt.Println("----------------------------------------------")

	std1 := test_data.Student{
		Id:5,
		Name:"小金李",
		Age:0,
		IsNewbie:true,
	}
	rset.Add(std1)	 // 添加一个元素
	printSet(rset)

	fmt.Println("----------------------------------------------")

	var newStudentCanInterface interface{} = []test_data.Student{}	// 新建一个装载结果的interface{}容器
	fmt.Println("fill前的newStudentCanInterface", newStudentCanInterface)

	rset.Fill(&newStudentCanInterface)								// 将rset填充到容器
	fmt.Println("fill之后的newStudentCanInterface", newStudentCanInterface)

	newStudentCan := newStudentCanInterface.([]test_data.Student)	// 将interface{}转换成其原始类型
	fmt.Println("强转对象之后的newStudentCan：", newStudentCan)

	test8()

	aa := []int{1,2,3,4}
	test1(aa)
	fmt.Println(aa)
	tt := aa[:2]
	fmt.Println("切片修改前", cap(tt))
	test1(tt)
	fmt.Println("切片修改后在方法外", cap(tt))
	fmt.Println(aa)

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

func test1(a []int) {
	a = append(a, 123)
	a[0] = 888
	fmt.Println("切片在方法中修改后", a)
}


func printSet(rset *collection.RSet) {
	for _, obj := range rset.Set {
		fmt.Printf("%d,%s,%d,%-v\n", obj["Id"], obj["Name"], obj["Age"], obj["IsNewbie"])
	}
}


func test8() {
	users := &[]user{{"liuliu", 20}, {"baobao", 21}, {"xiaowu", 22}}

	m := make(map[string]*user)
	for _, user := range *users {
		fmt.Println(user)
		m[user.Name] = &user
	}

	for name, u := range m {
		fmt.Println(name, u)
	}
}

type user struct {
	Name string
	Age  int
}