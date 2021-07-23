/**
 * string类型
 */

package main

import (
	"fmt"
	"unicode/utf8"
)

func stringTest1() {
	fmt.Println("stringTest1():")

	var str1 string = "abc"
	str2 := "def"
	fmt.Println(str1, str2)
	fmt.Printf("%T\n", str2)

	ch := 'a'
	str := "a" // "a" 实际上由 'a''\0' 组成
	fmt.Printf("%T, %T\n", ch, str)
}

func stringTest2() {
	fmt.Println("stringTest2():")

	/*
	 * 1. 使用内建函数len()获取字符串长度时, 字符串末尾的 \0 是不计算在内的
	 * 2. len()获取的是字符串中字节数, 而由于unicode使用可变字节长度, 一个中文占3字节, 所以 len(str2) == 6
	 */
	var str1 string = "abc"
	var str2 string = "你好"
	fmt.Println(len(str1))                    // 3
	fmt.Println(len(str2))                    // 6
	fmt.Println(utf8.RuneCountInString(str2)) // 2
}

func stringTest3() {
	fmt.Println("stringTest3():")

	str1 := "hello"
	str2 := "world"
	str3 := str1 + " " + str2 // 使用 + 进行字符串拼接
	fmt.Println(str3)
}

func stringTest4() {
	fmt.Println("stringTest4():")

	str := "hello world"
	fmt.Println(str[0]) // 可以通过下标去取字符

	for i, ch := range str {
		fmt.Printf("(%d, %c) ", i, ch)
	}
	fmt.Println()
}

func main() {
	stringTest1()
	stringTest2()
	stringTest3()
	stringTest4()
}
