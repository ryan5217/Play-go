package main

import (
	"Play-go/ryan"
	"fmt"
)

func main() {
	name := "Ryan"
	phone := 13162195208
	company := "上海宿七七酒店管理有限公司"
	fmt.Println("Hello world")
	fmt.Println(name)
	fmt.Println(phone)
	fmt.Println(company)

	classes := ryan.SayHello()
	fmt.Println(classes)
}
