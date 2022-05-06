package main

import (
	"fmt"
	"unsafe"
)

const (
	aa = "abc"
	bb = len(aa)
	cc = unsafe.Sizeof(aa)
)

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 15
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("Area is %d.", area)
	println()
	println(a, b, c)
	println(aa, bb, cc)
	println(Unknown, Female, Male)

	const (
		aaa = iota
		bbb
		ccc
		ddd = "ha"
		eee
		fff = 100
		ggg
		hhh = iota
		iii
	)
	fmt.Println(aaa, bbb, ccc, ddd, eee, fff, ggg, hhh, iii)

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	fmt.Println("l = ", l)

	println("---------- switch fallthrough ----------")
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
