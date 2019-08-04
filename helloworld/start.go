package main

import (
	"fmt"
	"math"
)

/**
go 语言不断学习中
*/
func main() {
	fmt.Println("hello go")
	startvar()
}

/**定义方法 初始化变量 */
func startvar() {
	var age int
	var name string
	//"%q"   空字符 “”
	fmt.Printf("%d %q\n", age, name)

	var i = 3
	var j = "hello"
	fmt.Println(i)
	fmt.Println(j)

	//var 定义变量可以在外面里面都行 ：=只能在函数里面
	a, b, f := 1, "s", true
	fmt.Printf("%d\n %s\n %d\n", a, b, f)

	//常量
	consts()
	//枚举
	enums()
}

func v() {
	//"内建变量使用"
	//bool
	//string
	//(u)int (u)int8  (u)int16 (u)int32 (u)int64   uintptr -->指针
	//byte rune   rune 就是go语言的char类型
	//float32 float64                 complex64 complex128 复数  实部 虚部
}

//go语言的常量定义
func consts() {
	//const hwl  ="licslan"
	//const a,b=3,4
	//上面也可以改为下面的方式  不要大写  大写有特殊含义注意
	const (
		hwl  = "licslan"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(hwl, c)
}

//go语言的枚举类型写法
func enums() {
	//const(
	//	cpp=1
	//	java=2
	//	c=3
	//	golang=5
	//)
	//上面的写法改为iota自增值
	const (
		cpp = iota
		_   //跳过不写
		java
		c
		golang
	)
	//定义 1024 自增  b kb mb gb tb pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Println(cpp, java, c, golang)
}
