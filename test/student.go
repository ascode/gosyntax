package main

type student struct {
	college
	name string
	age int
}

func main()  {
	//c := college{name:"huazhong tech"}
	s := student{
		college{name:"huazhong tech"},
		name: "jinfei",
		age: 33,
	}

	student{
		college,

	}

}