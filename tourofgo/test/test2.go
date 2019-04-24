package main

import (
	"fmt"
)

type Phone interface {
	call()
	music()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (nokiaPhone NokiaPhone) music() {
	fmt.Println("Nokia zzzzz...")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func (iPhone IPhone) music() {
	fmt.Println("IPhone zzzz...")
}

type HuaweiPhone struct {

}

func (huaweiPhone HuaweiPhone) call() {
	fmt.Println("I am huaweiPhonen, I can call you...")
}

func (huaweiPhone HuaweiPhone) music() {
	fmt.Println("I am huaweiPhonen, I can call you...")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
	phone.music()

	phone = new(IPhone)
	phone.call()

	phone = new(HuaweiPhone)
	phone.call()

}