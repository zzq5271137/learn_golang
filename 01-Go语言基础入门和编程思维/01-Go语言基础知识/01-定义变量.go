package main

import "fmt"

/**
 * Golang的变量定义
 */

func variableZeroValue() {
	/*
	 * go语言定义变量, 使用 var 关键字, 然后是变量名, 最后是变量类型
	 */
	var a int                   // int类型初始值默认为0
	var s string                // string类型初始值默认为空字符串
	fmt.Printf("%d %q\n", a, s) // %q: 字符串两边打印出引号
}

func variableInitialValue() {
	var a, b int = 3, 4 // 可以一次定义多个变量
	var s string = "abc"
	fmt.Println(a, b, s) // go语言中的变量一旦定义了, 就一定要用到, 否则报错
}

func variableTypeDeduction() {
	var a, b, s = 3, true, "abc" // 当定义变量并赋值时, 可以不声明变量类型, go语言编译器可以自动地决定类型
	fmt.Println(a, b, s)
}

func variableShorter() {
	a, b, s := 3, true, "abc" // 可以省略 var 关键字, 使用 := 定义变量并赋值 (这种方式只能在函数内使用)
	fmt.Println(a, b, s)
}

// aa := "outer"  // 在函数的外面定义变量不能使用 :=
var aa = "outer" // 这里定义的变量的作用域是在包内部 (第一行的 "package main")
var (            // 定义多个变量时可以使用这种方式
	bb = 2
	cc = true
	ss = "mine"
)

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc, ss)
}
