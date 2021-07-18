/**
 * string类型
 */

package main

import "fmt"

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
	 * 2. len()获取的是字符串中字符的个数, 一个字符为1字节, 而由于unicode使用可变字符长度, 一个中文占3字节, 所以 len(str2) == 6
	 */
	var str1 string = "abc"
	var str2 string = "你好"
	fmt.Println(len(str1)) // 3
	fmt.Println(len(str2)) // 6
}

func main() {
	stringTest1()
	stringTest2()
}
