package main

import (
	"fmt"
	"reflect"
)

func checkDatatype(val interface{}) reflect.Type {
	return reflect.TypeOf(val)
}

func funType() {
	panic("this is a function.")
}

func main() {
	// int
	var a int = 1
	fmt.Println("int:", a)

	var b float32 = 0.1
	fmt.Println("float:", b)

	var c complex64 = -3 + 4i
	fmt.Println("complex:", c)

	var d string = "sss" + `
	ddd
	`
	fmt.Println("string:", d)

	var e byte = 'a'
	fmt.Println("byte:", e)

	var f rune = '中'
	fmt.Println("rune:", f)

	var g bool = false
	fmt.Println("bool:", g)

	var h []int = []int{1, 3, 4, 2}
	fmt.Println("array:", h)

	h1 := [3]string{"1", "2", "3"}
	fmt.Println("array:", h1)

	h2 := [...]int{1, 2, 3}
	fmt.Println("array:", h2)

	var i []int
	fmt.Println("slice:", i)

	i1 := make([]int, 4)
	fmt.Println("slice:", i1)

	// map - var mapName map[keyType]valueType
	// map - make

	var j map[string]string = map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}
	var j1 = map[string]string{
		"d": "d",
	}
	fmt.Println("map:", j, j1)

	// struct

	type K struct {
		index int
		item  map[string]string
	}

	k := K{
		index: 1,
		item: map[string]string{
			"a": "a",
			"b": "b",
		},
	}
	fmt.Println("struct:", k)

	// interface
	type L interface {
		hello() string
	}

	var funcType = checkDatatype(funType)
	fmt.Println("func:", funcType)

	// %d 表示整型数字，%s 表示字符串

	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	// Code=123&endDate=2020-12-31
}
