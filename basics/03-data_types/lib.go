package main

import "fmt"

// 定义一个接口
type Animal interface {
	Sound() string
}

// 定义一个实现了Animal接口的类型
type Dog struct{}

// 实现Animal接口中的方法
func (d Dog) Sound() string {
	return "汪汪汪"
}

// 定义一个实现了Animal接口的类型
type Cat struct{}

// 实现Animal接口中的方法
func (c Cat) Sound() string {
	return "喵喵喵"
}

func T() {
	// 使用接口类型的变量来调用方法
	var animal Animal

	// 实例化一个Dog类型
	animal = Dog{}
	fmt.Println(animal.Sound()) // 输出：汪汪汪

	// 实例化一个Cat类型
	animal = Cat{}
	fmt.Println(animal.Sound()) // 输出：喵喵喵
}
