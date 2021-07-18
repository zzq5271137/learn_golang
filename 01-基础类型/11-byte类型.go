/**
 * byte类型
 */

package main

import "fmt"

func byteTest() {
	fmt.Println("byteTest():")

	var a byte = 'a' // golang中, 一个字符使用单引号, 字符串使用双引号
	fmt.Println(a)   // 97
	fmt.Printf("%c\n", a)
	fmt.Printf("%c\n", 97)
	fmt.Printf("%T\n", a) // uint8, byte类型是uint8的别名
	fmt.Println('\n')     // 10

	b := 'a'              // 这样进行自动类型推导, 实际上变量被声明成了rune类型
	fmt.Printf("%T\n", b) // int32, rune类型是int32的别名
	c := '你'
	fmt.Printf("%T\n", c)

	// 字符和字符串除了单双引号的区别外, 还有一个区别, 就是字符串的末尾都有一个 \0, 表示一个字符串结束的标志 (\0 对应的ASCII值为0)
}

func main() {
	byteTest()
}
