package main

var c, python, java bool //var用于声明一个变量列表
var i, j int = 1, 2

func add(a, b int32) int32 {
	return (a + b);
}

func processAToAB(a int) (x, y int){//命名返回值
	x = a * 2
	y = a / 2
	return
}

